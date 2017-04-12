package distiller

import (
	"github.com/google/syzkaller/tools/syz-strace/domain"
	"github.com/google/syzkaller/prog"
	"sort"
	"fmt"
)

type Distiller interface {
	MinCover(domain.Seeds) []*prog.Prog
	Add(domain.Seeds)
	Contributes(domain.Seed, map[uint64]bool) int
}

type DefaultDistiller struct {
	DistilledProgs []*prog.Prog
	CallToSeed map[*prog.Call]*domain.Seed
	CallToDistilledProg map[*prog.Call]*prog.Prog
	CallToIdx map[*prog.Call]int
	SeedDependencyGraph map[*domain.Seed]map[int]map[*prog.Arg]*prog.Arg
}

func NewDefaultDistiller() (d *DefaultDistiller) {
	d = &DefaultDistiller{
		DistilledProgs: make([]*prog.Prog, 0),
		CallToSeed: make(map[*prog.Call]*domain.Seed, 0),
		CallToDistilledProg: make(map[*prog.Call]*prog.Prog, 0),
		CallToIdx: make(map[*prog.Call]int, 0),
		SeedDependencyGraph: make(map[*domain.Seed]map[int]map[*prog.Arg]*prog.Arg, 0),
	}
	return
}

func (d *DefaultDistiller) Add(seeds domain.Seeds) {
	for _, seed := range seeds {
		d.CallToSeed[seed.Call] = seed
		d.SeedDependencyGraph[seed] = make(map[int]map[*prog.Arg]*prog.Arg, 0)
		seed.ArgMeta = make(map[*prog.Arg]bool, 0)
		for call, idx := range seed.DependsOn {
			if _, ok := d.SeedDependencyGraph[seed][idx]; !ok {
				d.SeedDependencyGraph[seed][idx] = make(map[*prog.Arg]*prog.Arg, 0)
			}
			d.CallToIdx[call] = idx
		}
		d.CallToIdx[seed.Call] = seed.CallIdx
	}
}

func (d *DefaultDistiller) Contributes(seed *domain.Seed, seenIps map[uint64]bool) int {
	total := 0
	for _, ip := range seed.Cover {
		if _, ok := seenIps[ip]; !ok {
			seenIps[ip] = true
			total += 1
		}
	}
	return total
}

func (d *DefaultDistiller) MinCover(seeds domain.Seeds) (distilled []*prog.Prog) {
	fmt.Printf("Computing Min Cover with %d seeds\n", len(seeds))
	seenIps := make(map[uint64]bool)
	sort.Sort(sort.Reverse(seeds))
	contributing_progs := 0
	for _, seed := range seeds {
		var ips int = d.Contributes(seed, seenIps)
		if ips > 0 {
			d.AddToDistilledProg(seed)
			fmt.Printf("Seed: %s contributes: %d ips out of its total of: %d\n", seed.Call.Meta.Name, ips, len(seed.Cover))
			contributing_progs += 1
		}
	}
	distilledProgs := make(map[*prog.Prog]bool)
	for _, seed := range seeds {
		if _, ok := d.CallToDistilledProg[seed.Call]; ok {
			distilledProgs[d.CallToDistilledProg[seed.Call]] = true
		}
	}
	for prog, _ := range distilledProgs {
		fmt.Printf("Prog: %v\n", prog)
		distilled = append(distilled, prog)
	}
	fmt.Printf("Total Contributing: %d, out of %d", contributing_progs, len(seeds))
	return
}

func (d *DefaultDistiller) TrackDependencies(prg *prog.Prog) {
	args := make(map[*prog.Arg]int, 0)
	for i, call := range prg.Calls {
		var seed *domain.Seed
		var ok bool
		if seed, ok = d.CallToSeed[call]; !ok {
			//Most likely an mmap we had to do
			fmt.Printf("Call: %s\n", call.Meta.CallName)
			continue
		}
		for _, arg := range call.Args {
			fmt.Printf("Arg: %s, %v\n", call.Meta.CallName, arg)
			upstream_maps := d.isDependent(arg, i, args)
			for k, argMap := range upstream_maps {
				fmt.Printf("K: %d\n", k)
				if d.SeedDependencyGraph[seed][k] == nil {
					d.SeedDependencyGraph[seed][k] = make(map[*prog.Arg]*prog.Arg, 0)
				}
				for argK, argV := range argMap {
					d.SeedDependencyGraph[seed][k][argK] = argV
				}
			}
		}
		fmt.Printf("depends on: %v\n", d.SeedDependencyGraph[seed])
		if call.Ret != nil {
			args[call.Ret] = i
		}
	}
}

func (d *DefaultDistiller) GetAllUpstreamDependents(seed *domain.Seed) []*prog.Call {
	calls := make([]*prog.Call, 0)
	callMap := make(map[*prog.Call]bool, 0)
	for idx, _ := range d.SeedDependencyGraph[seed] {
		call := seed.Prog.Calls[idx]
		if s, ok := d.CallToSeed[call]; ok {
			calls = append(calls, call)
			calls = append(calls, d.GetAllUpstreamDependents(s)...)
		} else {
			calls = append(calls, call)
		}
	}
	for _, call := range calls {
		callMap[call] = true
	}
	calls = make([]*prog.Call, 0)
	for k, _ := range callMap {
		calls = append(calls, k)
	}
	return calls
}

func (d *DefaultDistiller) AddToDistilledProg(seed *domain.Seed) {
	distilledProg := new(prog.Prog)
	distilledProg.Calls = make([]*prog.Call, 0)
	callIndexes := make([]int, 0)
	totalCalls := make([]*prog.Call, 0)

	if d.CallToDistilledProg[seed.Call] != nil {
		return
	}
	upstreamCalls := d.GetAllUpstreamDependents(seed)
	distinctProgs := d.getAllProgs(upstreamCalls)
	if len(distinctProgs) > 0 {
		totalCalls = d.getCalls(distinctProgs)
	} else {
		totalCalls = upstreamCalls
	}
	
	for _, call := range totalCalls {
		callIndexes = append(callIndexes, d.CallToIdx[call])
	}
	callIndexes = append(callIndexes, seed.CallIdx)
	sort.Ints(callIndexes)
	fmt.Printf("Call IDX: %v\n", callIndexes)
	for _, idx := range callIndexes {
		call := seed.Prog.Calls[idx]
		d.CallToDistilledProg[call] = distilledProg
		distilledProg.Calls = append(distilledProg.Calls, call)
	}
	for _, call := range distilledProg.Calls {
		if s, ok := d.CallToSeed[call]; ok {
			dependencyMap := d.SeedDependencyGraph[s]
			for idx, argMap := range dependencyMap {
				upstreamSeed := d.CallToSeed[seed.Prog.Calls[idx]]
				for argK, argV := range argMap {
					if _, ok := upstreamSeed.ArgMeta[argK]; !ok {
						fmt.Printf("UpstreamedSeed: %s, index: %d\n", upstreamSeed.Call.Meta.CallName, idx)
						argK.Uses = nil
						upstreamSeed.ArgMeta[argK] = true
					}
					if argK.Uses == nil {
						fmt.Printf("Allocating Uses: %s, index: %d\n", upstreamSeed.Call.Meta.CallName, idx)
						argK.Uses = make(map[*prog.Arg]bool, 0)
					}
					fmt.Printf("Setting ArgV: %s, %d\n", upstreamSeed.Call.Meta.CallName, idx) 
					argK.Uses[argV] = true
				}
			}
		}
	}
	seed.Call.Ret.Uses = nil
}

func (d *DefaultDistiller) getAllProgs(calls []*prog.Call) (ret []*prog.Prog) {
	distinctProgs := make(map[*prog.Prog]bool)
	for _, call := range calls {
		if _, ok := d.CallToDistilledProg[call]; ok {
			distinctProgs[d.CallToDistilledProg[call]] = true
		}
	}
	for k, _ := range distinctProgs {
		ret = append(ret, k)
	}
	return
}

func (d *DefaultDistiller) getCalls(progs []*prog.Prog) (ret []*prog.Call) {
	for _, p := range progs {
		ret = append(ret, p.Calls...)
	}
	return
}


func (d *DefaultDistiller) Clean(progDistilled *prog.Prog) {

}

func (d *DefaultDistiller) isDependent(arg *prog.Arg, callIdx int, args map[*prog.Arg]int) map[int]map[*prog.Arg]*prog.Arg {
	upstreamSet := make(map[int]map[*prog.Arg]*prog.Arg, 0)
	if arg == nil {
		return nil
	}
	//May need to support more kinds
	switch arg.Kind {
	case prog.ArgResult:
		//fmt.Printf("%v\n", args[arg.Res])
		if _, ok := args[arg.Res]; ok {
			if upstreamSet[args[arg.Res]] == nil {
				upstreamSet[args[arg.Res]] = make(map[*prog.Arg]*prog.Arg, 0)
			}
			upstreamSet[args[arg.Res]][arg.Res] = arg
		}
	case prog.ArgPointer:
		if _, ok := args[arg.Res]; ok {
			upstreamSet[args[arg.Res]][arg.Res] = arg
		} else {
			for k, argMap := range d.isDependent(arg.Res, callIdx, args) {
				if upstreamSet[k] == nil {
					upstreamSet[k] = make(map[*prog.Arg]*prog.Arg)
					upstreamSet[k] = argMap
				} else {
					for argK, argV := range argMap {
						upstreamSet[k][argK] = argV
					}
				}
			}
		}

	case prog.ArgGroup:
		for _, inner_arg := range arg.Inner {
			for k, argMap := range d.isDependent(inner_arg, callIdx, args) {
				if upstreamSet[k] == nil {
					upstreamSet[k] = make(map[*prog.Arg]*prog.Arg)
					upstreamSet[k] = argMap
				} else {
					for argK, argV := range argMap {
						upstreamSet[k][argK] = argV
					}
				}
			}
		}
	}
	args[arg] = callIdx
	//doesn't hurt to add again if it was already added
	return upstreamSet
}

