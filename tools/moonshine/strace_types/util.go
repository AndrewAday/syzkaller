package strace_types

import (
	"github.com/google/syzkaller/prog"
)

type Pair struct {
	A string
	B string
}

type SocketDesc struct {
	Domain string
	Type string
	Proto string
}

var (
	EnabledSyscalls = map[string]bool{}

	Unsupported = map[string]bool{
		"brk": true,
		//"mprotect": true,
		//"munmap": true,
		"": true,
		"execve": true, // unsupported
		"access": true, // unsupported
		//"mmap": true, // don't need, we generate our own
		//		"sendmsg": true, //TODO: the addr arg in msg_name struct is all wonky and ordering of args is off
		//		"recvmsg": true, //TODO: the addr arg in msg_name struct is all wonky and ordering of args is off
		"gettimeofday": true, // unsupported
		"kill": true, // unsupported
		//"keyctl": true,
		//"shmctl": true,
		//"getsockname": true,
		"arch_prctl": true,
		//"mremap": true, // knowing vma location is difficult
		"getcwd": true, // unsupported
		"setdomainname": true, // unsupported
		"reboot": true, // unsupported
		"getppid": true, // unsupported
		"umask": true, // unsupported
		"adjtimex": true, // unsupported
		//"ioctl$FIONBIO": true, // unsupported
		"rt_sigprocmask": true,
		"rt_sigaction": true,
		"rt_sigtimedwait":true,
		"sysfs": true, // unsupported
		"chdir": true, // unsupported
		"clone": true, // unsupported
		"newfstatat": true, // unsupported
		"getsid": true,
		"getcpu": true,
		"sched_get_priority_min": true,
		"sched_get_priority_max": true,
	}

	VMACall = map[string] bool {
		"mmap": true,
		"munmap": true,
		"mremap": true,
		"msync": true,
		"mprotect": true,
		"remap_file_pages": true,
		"shmat": true,
		"mlock": true,
		"munlock": true,
		"madvise": true,
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
		"sock_netlink": "$inet6",
	}

	Bind_labels = map[string]string {
		"fd": "",
		"sock": "",
		"sock_alg": "$alg",
		"sock_bt_hci": "$bt_hci",
		"sock_bt_l2cap": "$bt_l2cap",
		"sock_bt_rfcomm": "$bt_rfcomm",
		"sock_bt_sco": "$bt_sco",
		"sock_in": "$inet",
		"sock_in6": "$inet6",
		"sock_netlink": "$netlink",
		"sock_netrom": "$netrom",
		"sock_nfc_llcp": "$nfc_llcp",
		"sock_sctp": "$sctp",
		"sock_unix": "$unix",
		"AF_INET": "$inet",
		"AF_INET6": "$inet6",
		"sock_packet": "$packet",
	}

	Connect_labels = map[string]string {
		"fd": "",
		"sock": "",
		"sock_bt_l2cap": "$bt_l2cap",
		"sock_bt_rfcomm": "$bt_rfcomm",
		"sock_bt_sco": "$bt_sco",
		"sock_in": "$inet",
		"sock_in6": "$inet6",
		"sock_netlink": "$netlink",
		"sock_netrom": "$netrom",
		"sock_nfc_llcp": "$nfc_llcp",
		"sock_nfc_raw": "$nfc_raw",
		"sock_sctp": "$sctp",
		"sock_unix": "$unix",
		"AF_INET": "$inet",
		"AF_INET6": "$inet6",
	}

	Bpf_labels = map[string]string {
		"BPF_PROG_LOAD": "$PROG_LOAD",
	}

	Setsockopt_labels = map[Pair]string {
		Pair{"SOL_RAW", "0x7"}: "$sock_int",
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
		Pair{"IPPROTO_IPV6", "IPV6_RECVPKTINFO"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_RECVHOPLIMIT"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_RECVRTHDR"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_RECVHOPOPTS"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_RECVDSTOPTS"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_RECVTCLASS"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_2292HOPOPTS"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_2292HOPLIMIT"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_2292RTHDR"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_2292DSTOPTS"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_2292PKTINFO"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_RECVERR"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_UNICAST_HOPS"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_2292PKTOPTIONS"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_ADD_MEMBERSHIP"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_DROP_MEMBERSHIP"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_JOIN_ANYCAST"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_LEAVE_ANYCAST"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_FLOWLABEL_MGR"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_IPSEC_POLICY"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_XFRM_POLICY"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "MCAST_JOIN_GROUP"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "MCAST_BLOCK_SOURCE"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "MCAST_UNBLOCK_SOURCE"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "MCAST_LEAVE_GROUP"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "MCAST_JOIN_SOURCE_GROUP"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "MCAST_LEAVE_SOURCE_GROUP"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "MCAST_MSFILTER"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_PKTINFO"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_HOPOPTS"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_RTHDRDSTOPTS"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_RTHDR"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_DSTOPTS"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_PATHMTU"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IP6T_SO_GET_REVISION_MATCH"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IP6T_SO_GET_REVISION_TARGET"}: "$inet6_buf",
		Pair{"IPPROTO_TCP", "TCP_FASTOPEN"}: "$inet_tcp_int",
		Pair{"IPPROTO_IPV6", "IPV6_MTU_DISCOVER"}: "$inet6_mtu",
		Pair{"IPPROTO_IPV6", "IPV6_MTU"}: "$inet6_int",
		Pair{"SOL_IPV6", "IPV6_MTU_DISCOVER"}: "$inet6_mtu",
		Pair{"SOL_ICMPV6", "1"}: "$inet6_buf",
		Pair{"SOL_IPV6", "IPV6_MTU_DISCOVER"}: "$inet6_mtu",
		Pair{"SOL_PACKET", "PACKET_RX_RING"}: "$packet_rx_ring",
		Pair{"SOL_PACKET", "PACKET_RECV_OUTPUT"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_COPY_THRESH"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_AUXDATA"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_ORIGDEV"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_VERSION"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_HDRLEN"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_RESERVE"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_LOSS"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_VNET_HDR"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_TX_TIMESTAMP"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_TIMESTAMP"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_FANOUT"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_TX_HAS_OFF"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_QDISC_BYPASS"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_AUXDATA"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_HDRLEN"}: "$packet_rx_ring",
		Pair{"SOL_PACKET", "PACKET_ADD_MEMBERSHIP"}: "$packet_buf",
		Pair{"SOL_PACKET", "PACKET_DROP_MEMBERSHIP"}: "$packet_buf",
		Pair{"SOL_PACKET", "PACKET_STATISTICS"}: "$packet_buf",
		Pair{"SOL_PACKET", "PACKET_TX_RING"}: "$packet_tx_ring",
		Pair{"SOL_PACKET", "PACKET_FANOUT_DATA"}: "$packet_buf",
		Pair{"IPPROTO_IP", "IP_RECVERR"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_TOS"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_HDRINCL"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_ROUTER_ALERT"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_RECVOPTS"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_RETOPTS"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_PKTINFO"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_MTU_DISCOVER"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_RECVTTL"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_RECVTOS"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_MTU"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_FREEBIND"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_PASSEC"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_TRANSPARENT"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_RECVORIGDSTADDR"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_MINTTL"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_NODEFRAG"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_CHECKSUM"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_BIND_ADDRESS_NO_PORT"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_MULTICAST_TTL"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_MULTICAST_LOOP"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_MULTICAST_ALL"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_UNICAST_IF"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_OPTIONS"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_PKTOPTIONS"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_XFRM_POLICY"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_MULTICAST_IF"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_ADD_MEMBERSHIP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_DROP_MEMBERSHIP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_UNBLOCK_SOURCE"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_BLOCK_SOURCE"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_ADD_SOURCE_MEMBERSHIP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_DROP_SOURCE_MEMBERSHIP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_MSFILTER"}: "$inet_buf",
		Pair{"IPPROTO_IP", "MCAST_JOIN_GROUP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "MCAST_BLOCK_SOURCE"}: "$inet_buf",
		Pair{"IPPROTO_IP", "MCAST_UNBLOCK_SOURCE"}: "$inet_buf",
		Pair{"IPPROTO_IP", "MCAST_LEAVE_GROUP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "MCAST_JOIN_SOURCE_GROUP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "MCAST_LEAVE_SOURCE_GROUP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "MCAST_MS_FILTER"}: "$inet_buf",
		Pair{"IPPROTO_ICMP", "ICMP_FILTER"}: "$inet_icmp_ICMP_FILTER",
		Pair{"SOL_SCTP", "SCTP_EVENTS"}: "$inet_sctp_SCTP_EVENTS",
		Pair{"SOL_SCTP", "SCTP_INITMSG"}: "$inet_sctp_SCTP_INITMSG",
		Pair{"SOL_SCTP", "SCTP_RTOINFO"}: "$inet_sctp_SCTP_RTOINFO",
		Pair{"SOL_SCTP", "SCTP_AUTOCLOSE"}: "$inet_sctp_SCTP_AUTOCLOSE",
		Pair{"SOL_SCTP", "SCTP_PRIMARY_ADDR"}: "$inet_sctp_SCTP_PRIMARY_ADDR",
		Pair{"SOL_SCTP", "SCTP_ASSOCINFO"}: "$inet_sctp_SCTP_ASSOCINFO",
		Pair{"SOL_SCTP", "SCTP_PEER_ADDR_PARAMS"}: "$inet_sctp_SCTP_PEER_ADDR_PARAMS",
		Pair{"SOL_SCTP", "SCTP_DELAYED_ACK"}: "$inet_sctp_SCTP_DELAYED_ACK",
		Pair{"SOL_SCTP", "SCTP_DEFAULT_SEND_PARAM"}: "$inet_sctp_SCTP_DEFAULT_SEND_PARAM",
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
		Pair{"IPPROTO_IPV6", "IPV6_RECVPKTINFO"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_RECVHOPLIMIT"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_RECVRTHDR"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_RECVHOPOPTS"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_RECVDSTOPTS"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_RECVTCLASS"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_2292HOPOPTS"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_2292HOPLIMIT"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_2292RTHDR"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_2292DSTOPTS"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_2292PKTINFO"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_UNICAST_HOPS"}: "$inet6_int",
		Pair{"IPPROTO_IPV6", "IPV6_2292PKTOPTIONS"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_ADD_MEMBERSHIP"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_DROP_MEMBERSHIP"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_JOIN_ANYCAST"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_LEAVE_ANYCAST"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_FLOWLABEL_MGR"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_IPSEC_POLICY"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_XFRM_POLICY"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "MCAST_JOIN_GROUP"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "MCAST_BLOCK_SOURCE"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "MCAST_UNBLOCK_SOURCE"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "MCAST_LEAVE_GROUP"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "MCAST_JOIN_SOURCE_GROUP"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "MCAST_LEAVE_SOURCE_GROUP"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "MCAST_MSFILTER"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_PKTINFO"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_HOPOPTS"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_RTHDRDSTOPTS"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_RTHDR"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_DSTOPTS"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_PATHMTU"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IP6T_SO_GET_REVISION_MATCH"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IP6T_SO_GET_REVISION_TARGET"}: "$inet6_buf",
		Pair{"IPPROTO_IPV6", "IPV6_MTU_DISCOVER"}: "$inet6_mtu",
		Pair{"IPPROTO_IPV6", "IPV6_MTU"}: "$inet6_int",
		Pair{"SOL_IPV6", "IPV6_MTU_DISCOVER"}: "$inet6_mtu",
		Pair{"SOL_PACKET", "PACKET_RX_RING"}: "$packet_buf",
		Pair{"SOL_PACKET", "PACKET_RECV_OUTPUT"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_COPY_THRESH"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_AUXDATA"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_ORIGDEV"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_VERSION"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_HDRLEN"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_RESERVE"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_LOSS"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_VNET_HDR"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_TX_TIMESTAMP"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_TIMESTAMP"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_FANOUT"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_TX_HAS_OFF"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_QDISC_BYPASS"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_AUXDATA"}: "$packet_int",
		Pair{"SOL_PACKET", "PACKET_ADD_MEMBERSHIP"}: "$packet_buf",
		Pair{"SOL_PACKET", "PACKET_DROP_MEMBERSHIP"}: "$packet_buf",
		Pair{"SOL_PACKET", "PACKET_STATISTICS"}: "$packet_buf",
		Pair{"SOL_PACKET", "PACKET_TX_RING"}: "$packet_buf",
		Pair{"SOL_PACKET", "PACKET_FANOUT_DATA"}: "$packet_buf",
		Pair{"IPPROTO_IP", "IP_RECVERRR"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_TOS"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_HDRINCL"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_ROUTER_ALERT"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_RECVOPTS"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_RETOPTS"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_PKTINFO"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_MTU_DISCOVER"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_RECVTTL"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_RECVTOS"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_MTU"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_FREEBIND"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_PASSEC"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_TRANSPARENT"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_RECVORIGDSTADDR"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_MINTTL"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_NODEFRAG"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_CHECKSUM"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_BIND_ADDRESS_NO_PORT"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_MULTICAST_TTL"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_MULTICAST_LOOP"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_MULTICAST_ALL"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_UNICAST_IF"}: "$inet_int",
		Pair{"IPPROTO_IP", "IP_OPTIONS"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_PKTOPTIONS"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_XFRM_POLICY"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_MULTICAST_IF"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_ADD_MEMBERSHIP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_DROP_MEMBERSHIP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_UNBLOCK_SOURCE"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_BLOCK_SOURCE"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_ADD_SOURCE_MEMBERSHIP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_DROP_SOURCE_MEMBERSHIP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "IP_MSFILTER"}: "$inet_buf",
		Pair{"IPPROTO_IP", "MCAST_JOIN_GROUP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "MCAST_BLOCK_SOURCE"}: "$inet_buf",
		Pair{"IPPROTO_IP", "MCAST_UNBLOCK_SOURCE"}: "$inet_buf",
		Pair{"IPPROTO_IP", "MCAST_LEAVE_GROUP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "MCAST_JOIN_SOURCE_GROUP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "MCAST_LEAVE_SOURCE_GROUP"}: "$inet_buf",
		Pair{"IPPROTO_IP", "MCAST_MS_FILTER"}: "$inet_buf",
		Pair{"IPPROTO_ICMP", "ICMP_FILTER"}: "$inet_icmp_ICMP_FILTER",
		Pair{"SOL_SCTP", "SCTP_EVENTS"}: "$inet_sctp_SCTP_EVENTS",
		Pair{"SOL_SCTP", "SCTP_INITMSG"}: "$inet_sctp_SCTP_INITMSG",
		Pair{"SOL_SCTP", "SCTP_RTOINFO"}: "$inet_sctp_SCTP_RTOINFO",
		Pair{"SOL_SCTP", "SCTP_AUTOCLOSE"}: "$inet_sctp_SCTP_AUTOCLOSE",
		Pair{"SOL_SCTP", "SCTP_PRIMARY_ADDR"}: "$inet_sctp_SCTP_PRIMARY_ADDR",
		Pair{"SOL_SCTP", "SCTP_ASSOCINFO"}: "$inet_sctp_SCTP_ASSOCINFO",
		Pair{"SOL_SCTP", "SCTP_STATUS"}: "$inet_sctp_SCTP_STATUS",
		Pair{"SOL_SCTP", "SCTP_GET_PEER_ADDRS"}: "$inet_sctp_SCTP_GET_PEER_ADDRS",
		Pair{"SOL_SCTP", "SCTP_GET_LOCAL_ADDRS"}: "$inet_sctp_SCTP_GET_LOCAL_ADDRS",
		Pair{"SOL_SCTP", "SCTP_SOCKOPT_PEELOFF"}: "$inet_sctp_SCTP_SOCKOPT_PEELOFF",
		Pair{"SOL_SCTP", "SCTP_SOCKOPT_CONNECTX3"}: "$inet_sctp_SCTP_SOCOPT_CONNECTX3",
		Pair{"SOL_SCTP", "SCTP_PEER_ADDR_PARAMS"}: "$inet_sctp_SCTP_PEER_ADDR_PARAMS",
		Pair{"SOL_SCTP", "SCTP_DELAYED_ACK"}: "$inet_sctp_SCTP_DELAYED_ACK",
	}

	Getsockname_labels = map[string]string {
		"fd": "", // TODO: this is an illegal value. how do we interpret the uniontype?
		"sock": "",
		"sock_in": "$inet",
		"sock_in6": "$inet6",
		"sock_netlink": "$netlink",
		"sock_netrom": "$netrom",
		"sock_sctp": "$sctp",
		"sock_unix": "$unix",
	}

	SocketLevel_map = map[string]string{
		"SOL_SOCKET": "SOL_SOCKET",
		"SOL_IPV6": "IPPROTO_IPV6",
		"SOL_ICMPV6": "SOL_ICMPV6",
		"SOL_TCP": "IPPROTO_TCP",
		"SOL_RAW": "IPPROTO_ICMP",
		"SOL_PACKET": "SOL_PACKET",
		"SOL_IP": "IPPROTO_IP",
		"SOL_SCTP": "SOL_SCTP",
	}

	Sendto_labels = map[string]string {
		"fd": "",
		"sock": "",
		"sock_in": "$inet",
		"sock_in6": "$inet6",
		"sock_sctp": "$sctp",
		"sock_unix": "$unix",
		"sock_netlink": "$inet6",
	}

	Sendmsg_labels = map[string]string {
		"fd": "",
		"sock": "",
		"sock_in6": "",
		"sock_in": "",
		"sock_sctp": "$sctp",
		"sock_unix": "$unix",
		"sock_algconn": "$alg",
		"sock_kcm": "$kcm",
		"sock_netlink": "$netlink",
		"sock_netrom": "$netrom",
		"sock_nfc_llcp": "$nfc_llcp",
	}

	Recvfrom_labels = map[string]string {
		"fd": "",
		"sock": "",
		"sock_in": "$inet",
		"sock_in6": "$inet6",
		"sock_sctp": "$sctp",
		"sock_unix": "$unix",
		"sock_netlink": "$inet6",
		"sock_packet": "$packet",
	}

	Ioctl_map = map[string]string {
		"FIONBIO": "$int_in",
		"FIOASYNC": "$int_in",
		"FS_IOC_GETFLAGS": "$int_out",
		"FS_IOC_SETFLAGS": "$int_in",
		"SIOCGIFINDEX": "$sock_SIOCGIFINDEX",
	}

	Socket_labels_pair = map[Pair]string {
		Pair{"AF_INET", "SOCK_STREAM"}: "$inet",
		Pair{"AF_INET", "SOCK_DGRAM"}: "$inet",
		Pair{"AF_INET", "SOCK_STREAM"}: "$inet",
		Pair{"AF_INET", "SOCK_STREAM"}: "$inet",
		Pair{"AF_INET", "SOCK_STREAM"}: "$inet",
		Pair{"AF_INET", "SOCK_STREAM"}: "$inet",
		Pair{"AF_INET", "SOCK_STREAM"}: "$inet",
		Pair{"AF_INET", "SOCK_STREAM"}: "$inet",
		Pair{"AF_INET", "SOCK_STREAM"}: "$inet",
		Pair{"AF_INET", "SOCK_STREAM"}: "$inet",
		Pair{"AF_INET", "SOCK_RAW"}: "$inet_icmp_raw",
		Pair{"AF_PACKET", "SOCK_RAW"}: "$packet",
		Pair{"AF_PACKET", "SOCK_DGRAM"}: "$packet",
		Pair{"AF_INET", "SOL_SCTP"}: "$inet_sctp",
	}
	Socket_labels = map[string]string {
		"AF_INET": "$inet",
		"AF_INET6": "$inet6",
		"AF_KCM": "$kcm",
		"AF_UNIX": "$unix",
		"AF_NETLINK": "$netlink",
		"AF_PACKET": "$packet",
	}

	Fcntl_labels = map[string]string {
		"F_DUPFD": "$dupfd",
		"F_DUPFD_CLOEXEC": "$dupfd",
		"F_GETFD": "$getflags",
		"F_GETFL": "$getflags",
		"F_GETSIG": "$getflags",
		"F_GETLEASE": "$getflags",
		"F_GETPIPE_SZ": "$getflags",
		"F_GET_SEALS": "$getflags",
		"F_SETFD": "$setflags",
		"F_SETFL": "$setstatus",
		"F_SETLK": "$lock",
		"F_SETLKW": "$lock",
		"F_GETLK": "$lock",
		"F_GETOWN": "$getown",
		"F_SETOWN": "$setown",
		"F_GETOWN_EX": "$getownex",
		"F_SETOWN_EX": "$setownex",
		"F_SETSIG": "$setsig",
		"F_SETLEASE": "$setlease",
		"DN_MULTISHOT": "$notify",
		"DN_ACCESS": "$notify",
		"DN_MODIFY": "$notify",
		"DN_CREATE": "$notify",
		"DN_DELETE": "$notify",
		"DN_RENAME": "$notify",
		"DN_ATTRIB": "$notify",
		"F_SETPIPE_SZ": "$setpipe",
		"F_ADD_SEALS": "$addseals",
	}

	Keyctl_labels = map[string]string {
		"KEYCTL_GET_KEYRING_ID": "$get_keyring_id",
		"KEYCTL_JOIN_SESSION_KEYRING": "$join",
		"KEYCTL_UPDATE": "$update",
		"KEYCTL_REVOKE": "$revoke",
		"KEYCTL_DESCRIBE": "$describe",
		"KEYCTL_CLEAR": "$clear",
		"KEYCTL_LINK": "$link",
		"KEYCTL_UNLINK": "$unlink",
		"KEYCTL_SEARCH": "$search",
		"KEYCTL_READ": "$read",
		"KEYCTL_CHOWN": "$chown",
		"KEYCTL_SETPERM": "$setperm",
		"KEYCTL_INSTANTIATE": "$instantiate",
		"KEYCTL_NEGATE": "$negate",
		"KEYCTL_SET_REQKEY_KEYRING": "$set_reqkey_keyring",
		"KEYCTL_SET_TIMEOUT": "$set_timeout",
		"KEYCTL_ASSUME_AUTHORITY": "$assume_authority",
		"KEYCTL_GET_SECURITY": "$get_security",
		"KEYCTL_SESSION_TO_PARENT": "$session_to_parent",
		"KEYCTL_REJECT": "$reject",
		"KEYCTL_INSTANTIATE_IOV": "$instantiate_iov",
		"KEYCTL_INVALIDATE": "$invalidate",
		"KEYCTL_GET_PERSISTENT": "$get_persistent",
	}

	Prctl_labels = map[string]string {
		"PR_GET_DUMPABLE": "$void",
		"PR_GET_KEEPCAPS": "$void",
		"PR_GET_NO_NEW_PRIVS": "$void",
		"PR_GET_SECCOMP": "$void",
		"PR_GET_SECUREBITS": "$void",
		"PR_GET_TIMERSLACK": "$void",
		"PR_GET_TIMING": "$void",
		"PR_TASK_PERF_EVENTS_DISABLE": "$void",
		"PR_TASK_PERF_EVENTS_ENABLE": "$void",
		"PR_MCE_KILL_GE": "$void",
		"PR_CAPBSET_READ": "$intptr",
		"PR_CAPBSET_DROP": "$intptr",
		"PR_SET_CHILD_SUBREAPER": "$intptr",
		"PR_SET_DUMPABLE": "$intptr",
		"PR_SET_FPEMU": "$intptr",
		"PR_SET_KEEPCAPS": "$intptr",
		"PR_SET_NO_NEW_PRIVS": "$intptr",
		"PR_SET_PDEATHSIG": "$intptr",
		"PR_SET_SECUREBITS": "$intptr",
		"PR_SET_TIMERSLACK": "$intptr",
		"PR_SET_TIMING": "$intptr",
		"PR_SET_TSC": "$intptr",
		"PR_SET_UNALIGN": "$intptr",
		"PR_MCE_KILL": "$intptr",
		"PR_GET_CHILD_SUBREAPER": "$getreaper",
		"PR_GET_ENDIAN": "$getreaper",
		"PR_GET_FPEMU": "$getreaper",
		"PR_GET_FPEXC": "$getreaper",
		"PR_GET_PDEATHSIG": "$getreaper",
		"PR_GET_TID_ADDRESS": "$getreaper",
		"PR_GET_TSC": "$getreaper",
		"PR_GET_UNALIGN": "$getreaper",
		"PR_SET_ENDIAN": "$setendian",
		"PR_SET_FPEXC": "$setfpexc",
		"PR_SET_NAME": "$setname",
		"PR_GET_NAME": "$getname",
		"PR_SET_PTRACER": "$setptracer",
	}

	Special_Consts = map[string]uint64 {
		"_LINUX_CAPABILITY_VERSION_1": uint64(0x19980330),
		"_LINUX_CAPABILITY_VERSION_2": uint64(0x20071026),
		"_LINUX_CAPABILITY_VERSION_3": uint64(0x20080522),
		"PROT_NONE": 0,
		"O_ASYNC": uint64(0x0040),
		"O_TMPFILE": uint64(0x022000000),
		"SIGUSR1": uint64(10),
		"SIGALRM": uint64(14),
		"SIGFPE": uint64(8),
		"SIG_0": uint64(0),
		"S_ISUID": uint64(0x0004000),
		"S_ISGID": uint64(0x0002000),
		"RLIM64_INFINITY": ^uint64(0),
		"CLOCK_REALTIME_ALARM": uint64(8),
		"CLOCK_BOOTTIME_ALARM": uint64(9),
		"PERF_ATTR_SIZE_VER5": uint64(112),
		"PERF_ATTR_SIZE_VER4": uint64(104),
		"PERF_ATTR_SIZE_VER3": uint64(96),
		"PERF_ATTR_SIZE_VER2": uint64(80),
		"PERF_ATTR_SIZE_VER1": uint64(72),
		"PERF_ATTR_SIZE_VER0": uint64(64),
		"PERF_COUNT_HW_CPU_CYCLES": uint64(0),
		"PERF_COUNT_HW_INSTRUCTIONS": uint64(1),
		"PERF_COUNT_HW_CACHE_REFERENCES": uint64(2),
		"PERF_COUNT_HW_CACHE_MISSES": uint64(3),
		"PERF_COUNT_HW_BRANCH_INSTRUCTIONS": uint64(4),
		"PERF_COUNT_HW_BRANCH_MISSES": uint64(5),
		"PERF_COUNT_HW_BUS_CYCLES": uint64(6),
		"PERF_COUNT_HW_STALLED_CYCLES_FRONTEND" : uint64(7),
		"PERF_COUNT_HW_STALLED_CYCLES_BACKEND" : uint64(8),
		"PERF_COUNT_HW_REF_CPU_CYCLES": uint64(9),
		"PERF_COUNT_SW_CPU_CLOCK": uint64(0),
		"PERF_COUNT_SW_TASK_CLOCK": uint64(1),
		"PERF_COUNT_SW_PAGE_FAULTS": uint64(2),
		"PERF_COUNT_SW_CONTEXT_SWITCHES": uint64(3),
		"PERF_COUNT_SW_CPU_MIGRATIONS": uint64(4),
		"PERF_COUNT_SW_PAGE_FAULTS_MIN": uint64(5),
		"PERF_COUNT_SW_PAGE_FAULTS_MAJ": uint64(6),
		"PERF_COUNT_SW_ALIGNMENT_FAULTS": uint64(7),
		"PERF_COUNT_SW_EMULATION_FAULTS": uint64(8),
		"FUTEX_WAIT_PRIVATE": uint64(128),
		"FUTEX_WAKE_PRIVATE": uint64(129),
		"FUTEX_WAIT_REQUEUE_PI_PRIVATE":  uint64(139),
		"FUTEX_UNLOCK_PI_PRIVATE": uint64(135),
		"FUTEX_CMP_REQUEUE_PI_PRIVATE": uint64(140),
		"MCL_ONFAULT": uint64(32768),
		"CAP_CHOWN": uint64(0),
		"CAP_DAC_OVERRIDE": uint64(1),
		"CAP_DAC_READ_SEARCH": uint64(2),
		"CAP_FOWNER": uint64(3),
		"CAP_FSETID": uint64(4),
		"CAP_KILL": uint64(5),
		"CAP_SETGID": uint64(6),
		"CAP_SETUID": uint64(7),
		"CAP_SETPCAP": uint64(8),
		"CAP_LINUX_IMMUTABLE": uint64(9),
		"CAP_NET_BIND_SERVICE": uint64(10),
		"CAP_NET_BROADCAST": uint64(11),
		"CAP_NET_ADMIN": uint64(12),
		"CAP_NET_RAW": uint64(13),
		"CAP_IPC_LOCK": uint64(14),
		"CAP_IPC_OWNER": uint64(15),
		"CAP_SYS_MODULE": uint64(16),
		"CAP_SYS_RAWIO": uint64(17),
		"CAP_SYS_CHROOT": uint64(18),
		"CAP_SYS_PTRACE": uint64(19),
		"CAP_SYS_PACCT": uint64(20),
		"CAP_SYS_ADMIN": uint64(21),
		"CAP_SYS_BOOT" : uint64(22),
		"CAP_SYS_NICE" : uint64(23),
		"CAP_SYS_RESOURCE": uint64(24),
		"CAP_SYS_TIME": uint64(25),
		"CAP_SYS_TTY_CONFIG": uint64(26),
		"CAP_MKNOD": uint64(27),
		"CAP_LEASE": uint64(28),
		"CAP_AUDIT_WRITE": uint64(29),
		"CAP_AUDIT_CONTROL": uint64(30),
		"CAP_SETFCAP": uint64(31),
		"FALLOC_FL_COLLAPSE_RANGE": uint64(0x08),
		"FALLOC_FL_ZERO_RANGE": uint64(0x10),
		"FALLOC_FL_INSERT_RANGE": uint64(0x20),
		"FALLOC_FL_UNSHARE_RANGE": uint64(0x40),
		"MEMBARRIER_CMD_QUERY": uint64(0),
		"MEMBARRIER_CMD_GLOBAL": uint64(1 << 0),
		"MEMBARRIER_CMD_GLOBAL_EXPEDITED": uint64(1 << 1),
		"MEMBARRIER_CMD_REGISTER_GLOBAL_EXPEDITED": uint64(1 << 2),
		"MEMBARRIER_CMD_PRIVATE_EXPEDITED": uint64(1 << 3),
		"MEMBARRIER_CMD_REGISTER_PRIVATE_EXPEDITED": uint64(1 << 4),
		"MEMBARRIER_CMD_PRIVATE_EXPEDITED_SYNC_CORE": uint64(1 << 5),
		"MEMBARRIER_CMD_REGISTER_PRIVATE_EXPEDITED_SYNC_CORE": uint64(1 << 6),
		"MEMBARRIER_CMD_SHARED": uint64(1 << 0),
		"SIGUSR2": 31,
		"R_OK": uint64(4),
		"W_OK": uint64(2),
		"X_OK": uint64(1),
		"F_OK": uint64(0),
	}

)

func commonArg(t prog.Type) prog.ArgCommon {
	common := prog.ArgCommon{}
	common.AddType(t)
	return common
}

func GroupArg(t prog.Type, inner []prog.Arg) prog.Arg {
	return &prog.GroupArg{ArgCommon: commonArg(t), Inner: inner}
}

func PointerArg(t prog.Type, page uint64, off int, npages uint64, obj prog.Arg) prog.Arg {
	return &prog.PointerArg{ArgCommon: commonArg(t), PageIndex: page, PageOffset: off, PagesNum: npages, Res: obj}
}

func ConstArg(t prog.Type, v uint64) prog.Arg {
	return &prog.ConstArg{ArgCommon: commonArg(t), Val: v}
}

func DataArg(t prog.Type, data []byte) prog.Arg {
	return &prog.DataArg{ArgCommon: commonArg(t), Data: append([]byte{}, data...)}
}

func UnionArg(t prog.Type, opt prog.Arg, typ prog.Type) prog.Arg {
	return &prog.UnionArg{ArgCommon: commonArg(t), Option: opt, OptionType: typ}
}

func ResultArg(t prog.Type, r prog.Arg, v uint64) prog.Arg {
	arg := &prog.ResultArg{ArgCommon: commonArg(t), Res: r, Val: v}
	if r == nil {
		return arg
	}
	if used, ok := r.(prog.ArgUsed); ok {
		if *used.Used() == nil {
			*used.Used() = make(map[prog.Arg]bool)
		}
		if (*used.Used())[arg] {
			panic("already used")
		}
		(*used.Used())[arg] = true
	}
	return arg
}

func ReturnArg(t prog.Type) prog.Arg {
	return &prog.ReturnArg{ArgCommon: commonArg(t)}
}

func GetSyzType(typ prog.Type) string {
	switch a := typ.(type) {
	case *prog.ResourceType:
		return "ResourceType-" + a.Desc.Kind[0]
	case *prog.BufferType:
		return "BufferType"
	case *prog.VmaType:
		return "VmaType"
	case *prog.FlagsType:
		return "FlagsType"
	case *prog.ConstType:
		return "ConstType"
	case *prog.IntType:
		return "IntType"
	case *prog.ProcType:
		return "ProcType"
	case *prog.ArrayType:
		return "ArrayType"
	case *prog.StructType:
		return "StructType"
	case *prog.UnionType:
		return "UnionType"
	case *prog.PtrType:
		return "PtrType"
	case *prog.LenType:
		return "LenType"
	case *prog.CsumType:
		return "CsumType"
	default:
		panic("unknown argument type")
	}
}




