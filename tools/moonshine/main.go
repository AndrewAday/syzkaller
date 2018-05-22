package main

import (
	. "github.com/google/syzkaller/tools/moonshine/scanner"
	. "github.com/google/syzkaller/tools/moonshine/parser"
	"github.com/google/syzkaller/prog"
	"github.com/google/syzkaller/pkg/hash"
	"fmt"
	"os"
	"github.com/google/syzkaller/pkg/db"
	"io/ioutil"
	"path/filepath"
	"strings"
	"strconv"
	"flag"
	"github.com/google/syzkaller/tools/moonshine/strace_types"
	. "github.com/google/syzkaller/tools/moonshine/logging"
	"github.com/google/syzkaller/sys"
	"path"
)

var (
	flagFile = flag.String("file", "", "file to parse")
	flagDir = flag.String("dir", "", "director to parse")
)

const (
	OS = "linux"
	Arch = "amd64"
	currentDBVersion = 3

)

func main() {
	fmt.Printf("git revision: %s\n", sys.GitRevision)
	flag.Parse()
	target, err := prog.GetTarget(OS, Arch)
	if err != nil {
		Failf("error getting target: %v", err.Error())
	} else {
		fmt.Printf("Const Map len: %d\n", len(target.ConstMap))
		ParseTraces(target, false)
		pack("serialized", "corpus.db")
	}
}

func progIsTooLarge(prog_ *prog.Prog) bool {
	buff := make([]byte, prog.ExecBufferSize)
	if _, err := prog_.SerializeForExec(buff); err != nil {
		return true
	}
	return false
}

func ParseTraces(target *prog.Target, distill bool) []*prog.Prog {
	ret := make([]*prog.Prog, 0)
	names := make([]string, 0)
	if *flagFile != "" {
		names = append(names, *flagFile)
	} else if *flagDir != "" {
		names = getFileNames(*flagDir)
	} else {
		panic("Flag or FlagDir required")
	}
	for _, file := range(names) {
		fmt.Printf("Scanning file: %s\n", file)
		tree := Parse(file)
		if tree == nil {
			fmt.Fprintf(os.Stderr, "File: %s is empty\n", file)
			continue
		}
		progs := ParseTree(tree, tree.RootPid, target)
		ret = append(ret, progs...)
		i := 0
		for _, prog_ := range progs {
			prog_.Target = target
			if !distill {
				if progIsTooLarge(prog_) {
					fmt.Fprintln(os.Stderr, "Prog is too large\n")
					continue
				}
				i += 1
				s_name := "serialized/" + filepath.Base(file) + strconv.Itoa(i)
				if err := ioutil.WriteFile(s_name, prog_.Serialize(), 0640); err != nil {
					Failf("failed to output file: %v", err)
				}
			}
		}

	}
	return ret
}

func getFileNames(dir string) []string {
	names := make([]string, 0)
	if infos, err := ioutil.ReadDir(dir); err == nil {
		for _, info := range (infos) {
			name := path.Join(dir, info.Name())
			names = append(names, name)
		}
	} else {
		Failf("Failed to read dir: %s\n", err.Error())
	}
	return names
}

func ParseTree(tree *strace_types.TraceTree, pid int64, target *prog.Target) []*prog.Prog {
	fmt.Fprintf(os.Stderr, "Parsing tree for file: %s\n", tree.Filename)
	progs := make([]*prog.Prog, 0)
	parsedProg, ctx, err := ParseProg(tree.TraceMap[pid], target)
	if err != nil {
		panic("Failed to parse program")
	} else {
		if len(parsedProg.Calls) == 0 {
			parsedProg = nil
		}else {
			if err = ctx.State.Tracker.FillOutMemory(parsedProg); err != nil {
				fmt.Fprintf(os.Stderr, "Out of bounds memory: %s %d\n", tree.Filename, pid)
				parsedProg = nil
			} else {
				totalMemory := ctx.State.Tracker.GetTotalMemoryAllocations(parsedProg)
				if totalMemory == 0 {
					fmt.Printf("length of zero mem prog: %d\n", totalMemory)
				} else {
					mmapCall := ctx.Target.MakeMmap(0, uint64(totalMemory))
					calls := make([]*prog.Call, 0)
					calls = append(append(calls, mmapCall), parsedProg.Calls...)
					parsedProg.Calls = calls
				}
				if err := parsedProg.Validate(); err != nil {
					panic(fmt.Sprintf("Error validating program: %s\n", err.Error()))
				}
			}
		}
	}
	if parsedProg != nil {
		fmt.Fprintf(os.Stderr, "Appending program: %s %d\n", tree.Filename, pid)
		progs = append(progs, parsedProg)
	}
	for _, pid_ := range(tree.Ptree[pid]) {
		if tree.TraceMap[pid_] != nil{
			progs = append(progs, ParseTree(tree, pid_, target)...)
		}
	}
	return progs
}


func pack(dir, file string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		Failf("failed to read dir: %v", err)
	}
	os.Remove(file)
	db, err := db.Open(file)
	db.BumpVersion(currentDBVersion)
	if err != nil {
		Failf("failed to open database file: %v", err)
	}
	for _, file := range files {
		data, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			Failf("failed to read file %v: %v", file.Name(), err)
		}
		var seq uint64
		key := file.Name()
		if parts := strings.Split(file.Name(), "-"); len(parts) == 2 {
			var err error

			if seq, err = strconv.ParseUint(parts[1], 10, 64); err == nil {
				key = parts[0]
			}
		}
		if sig := hash.String(data); key != sig {
			fmt.Fprintf(os.Stdout, "fixing hash %v -> %v\n", key, sig)
			key = sig
		}
		db.Save(key, data, seq)
	}
	if err := db.Flush(); err != nil {
		Failf("failed to save database file: %v", err)
	}
}
