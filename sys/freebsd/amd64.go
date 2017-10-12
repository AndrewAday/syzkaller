// AUTOGENERATED FILE
package freebsd

import . "github.com/google/syzkaller/prog"

func init() {
	RegisterTarget(&Target{OS: "freebsd", Arch: "amd64", Revision: revision_amd64, PtrSize: 8, Syscalls: syscalls_amd64, Resources: resources_amd64, Structs: structDescs_amd64, Consts: consts_amd64}, initTarget)
}

var resources_amd64 = []*ResourceDesc(nil)

var structDescs_amd64 = []*KeyedStruct(nil)

var syscalls_amd64 = []*Syscall{
	{NR: 1, Name: "mmap", CallName: "mmap", Args: []Type{
		&VmaType{TypeCommon: TypeCommon{TypeName: "vma", FldName: "addr", TypeSize: 8}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "len", TypeSize: 8}}, Buf: "addr"},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "mmap_prot", FldName: "prot", TypeSize: 8}}, Vals: []uint64{2, 3}},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "mmap_flags", FldName: "flags", TypeSize: 8}}, Vals: []uint64{4, 5, 6}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "fd", TypeSize: 8}}, Val: 18446744073709551615},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "offset", TypeSize: 8}}},
	}, Ret: &VmaType{TypeCommon: TypeCommon{TypeName: "vma", FldName: "ret", TypeSize: 8, ArgDir: 1}}},
}

var consts_amd64 = []ConstValue{
	{Name: "MAP_ANONYMOUS", Value: 5},
	{Name: "MAP_FIXED", Value: 6},
	{Name: "MAP_PRIVATE", Value: 4},
	{Name: "PROT_READ", Value: 2},
	{Name: "PROT_WRITE", Value: 3},
	{Name: "SYS_mmap", Value: 1},
}

const revision_amd64 = "8e89c92ba6210bfbb714647598abe10b8cbb4edf"
