package syz_structs

type Stack []byte

func (s Stack) Push(c byte) Stack {
	return append(s, c)
}

func (s Stack) Pop() (Stack, byte) {
	l := len(s)
	if l == 0 {
		panic("popping an empty stack")
	}
	r := s[l-1]
	s = s[:l-1]
	return s, r
}

type Pair struct {
	A string
	B string
}

var (
	Unsupported = map[string]bool{
		"brk": true,
		//"fstat": true,
		//"exit_group": true,
		"mprotect": true,
		"munmap": true,
		"execve": true,
		"access": true,
		"mmap": true,
		//"accept": true, // need to determine accept type from the type of sockfd. Unsure how to do this cleanly.
		"bind": true, // same issue
		"sendto": true, // same
		// also: problem with select, 2nd arg not correct format
		"select": true,
		"recvfrom": true,
		//"socket": true, // ltp_asapi_03 has comment in format!!
		"sendmsg": true,
		"recvmsg": true,
		"gettimeofday": true,
		"getsockname": true,
		"connect": true,
		"getsockopt": true,
		//"accept4": true,
		"mremap": true, // knowing vma location is difficult
		"getcwd": true, // unsupported
		"setdomainname": true, // unsupported
		"reboot": true, // unsupported
		"getppid": true, // unsupported
		"umask": true, // unsupported
		"adjtimex": true, // unsupported
		"ioctl$FIONBIO": true, // unsupported
		"sysfs": true,
		"chdir": true,
		"fcntl": true,
		//"arch_prctl": true, // has two conflicting method signatures!! http://man7.org/linux/man-pages/man2/arch_prctl.2.html
		//"rt_sigaction": true, // constants such as SIGRTMIN are not defined in syzkaller, and missing last void __user *, restorer argument
		//"rt_sigprocmask": true, // second arg given as an array, should be pointer
		//"getrlimit": true, // has arg 8192*1024, cannot evaluate easily
		//"statfs": true, // types disagree, strace gives struct, syzkaller expects buffer
		//"fstatfs": true, // types disagree, strace gives struct, syzkaller expects buffer
		//"ioctl": true, // types disagree, strace gives struct, syzkaller expects buffer
		/* can build the ioctl$arg from the 2nd arg */
		//"getdents": true, // types disagree, strace gives struct, syzkaller expects buffer
	}

	Accept_labels = map[string]string {
		"fd": "", // TODO: this is an illegal value. how do we interpret the uniontype?
		"sock": "",
		"sock_alg": "$alg",
		"sock_in": "$inet",
		"sock_in6": "$inet6",
		"sock_netrom": "$netrom",
		"sock_nfc_llcp": "$nfc_llcp",
		"sock_sctp": "$sctp",
		"sock_unix": "$unix",
	}

	Setsockopt_labels = map[Pair]string {
		Pair{"SOL_SOCKET","SO_DETACH_FILTER"}: "$sock_void",
		Pair{"SOL_SOCKET","SO_MARK"}: "$sock_void",
		Pair{"SOL_SOCKET","SO_ACCEPTCONN"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_BROADCAST"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_DEBUG"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_DOMAIN"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_ERROR"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_DONTROUTE"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_KEEPALIVE"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_PEEK_OFF"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_PRIORITY"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_PROTOCOL"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_RCVBUF"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_RCVBUFFORCE"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_RCVLOWAT"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_SNDLOWAT"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_REUSEADDR"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_SNDBUF"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_SNDBUFFORCE"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_TIMESTAMP"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_TYPE"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_REUSEPORT"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_OOBINLINE"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_NO_CHECK"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_PASSCRED"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_TIMESTAMPNS"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_LOCK_FILTER"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_PASSSEC"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_RXQ_OVFL"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_WIFI_STATUS"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_NOFCS"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_SELECT_ERR_QUEUE"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_BUSY_POLL"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_MAX_PACING_RAT"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_BINDTODEVICE"}: "$sock_str",
		Pair{"SOL_SOCKET","SO_LINGER"}: "$sock_linger",
		Pair{"SOL_SOCKET","SO_PEERCRED"}: "$sock_cred",
		Pair{"SOL_SOCKET","SO_RCVTIMEO"}: "$sock_timeval",
		Pair{"SOL_SOCKET","SO_SNDTIMEO"}: "$sock_timeval",
		Pair{"SOL_SOCKET","SO_ATTACH_BPF"}: "$sock_attack_bpf",
		Pair{"SOL_SOCKET","SO_TIMESTAMPING"}: "$SO_TIMESTAMPING",
		Pair{"SOL_SOCKET","SO_ATTACH_FILTER"}: "$SO_ATTACH_FILTER",
	}

	Getsockopt_labels = map[Pair]string {
		Pair{"SOL_SOCKET","SO_ACCEPTCONN"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_BROADCAST"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_DEBUG"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_DOMAIN"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_ERROR"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_DONTROUTE"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_KEEPALIVE"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_PEEK_OFF"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_PRIORITY"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_PROTOCOL"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_RCVBUF"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_RCVBUFFORCE"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_RCVLOWAT"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_SNDLOWAT"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_REUSEADDR"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_SNDBUF"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_SNDBUFFORCE"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_TIMESTAMP"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_TYPE"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_REUSEPORT"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_OOBINLINE"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_NO_CHECK"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_PASSCRED"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_TIMESTAMPNS"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_LOCK_FILTER"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_PASSSEC"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_RXQ_OVFL"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_WIFI_STATUS"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_NOFCS"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_SELECT_ERR_QUEUE"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_BUSY_POLL"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_MAX_PACING_RAT"}: "$sock_int",
		Pair{"SOL_SOCKET","SO_LINGER"}: "$sock_linger",
		Pair{"SOL_SOCKET","SO_PEERCRED"}: "$sock_cred",
		Pair{"SOL_SOCKET","SO_RCVTIMEO"}: "$sock_timeval",
		Pair{"SOL_SOCKET","SO_SNDTIMEO"}: "$sock_timeval",
		Pair{"SOL_SOCKET","SO_TIMESTAMPING"}: "$SO_TIMESTAMPING",
		Pair{"SOL_SOCKET","SO_BINDTODEVICE"}: "$sock_buf",
		Pair{"SOL_SOCKET","SO_PEERNAME"}: "$sock_buf",
		Pair{"SOL_SOCKET","SO_PEERSEC"}: "$sock_buf",
		Pair{"SOL_SOCKET","SO_GET_FILTE"}: "$sock_buf",
	}

	Socket_labels = map[string]string {
		"AF_INET": "$inet",
		"AF_INET6": "$inet6",
		"AF_KCM": "$kcm",


	}
)