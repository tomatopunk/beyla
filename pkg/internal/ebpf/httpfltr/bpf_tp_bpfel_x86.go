// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package httpfltr

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_tpConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpf_tpHttpBufT struct {
	Flags    uint64
	ConnInfo bpf_tpConnectionInfoT
	Buf      [1024]uint8
	_        [4]byte
}

type bpf_tpHttpConnectionMetadataT struct {
	Pid struct {
		HostPid   uint32
		UserPid   uint32
		Namespace uint32
	}
	Type uint8
}

type bpf_tpHttpInfoT struct {
	Flags           uint64
	ConnInfo        bpf_tpConnectionInfoT
	_               [4]byte
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [160]uint8
	Len             uint32
	RespLen         uint32
	Status          uint16
	Type            uint8
	Ssl             uint8
	Pid             struct {
		HostPid   uint32
		UserPid   uint32
		Namespace uint32
	}
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
}

type bpf_tpPidConnectionInfoT struct {
	Conn bpf_tpConnectionInfoT
	Pid  uint32
}

type bpf_tpPidKeyT struct {
	Pid       uint32
	Namespace uint32
}

type bpf_tpRecvArgsT struct {
	SockPtr  uint64
	IovecPtr uint64
}

type bpf_tpSockArgsT struct {
	Addr       uint64
	AcceptTime uint64
}

type bpf_tpSslArgsT struct {
	Ssl    uint64
	Buf    uint64
	LenPtr uint64
}

type bpf_tpTpInfoPidT struct {
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	Pid   uint32
	Valid uint8
	_     [3]byte
}

// loadBpf_tp returns the embedded CollectionSpec for bpf_tp.
func loadBpf_tp() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_tpBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_tp: %w", err)
	}

	return spec, err
}

// loadBpf_tpObjects loads bpf_tp and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_tpObjects
//	*bpf_tpPrograms
//	*bpf_tpMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_tpObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_tp()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_tpSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tpSpecs struct {
	bpf_tpProgramSpecs
	bpf_tpMapSpecs
}

// bpf_tpSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tpProgramSpecs struct {
	KprobeSysExit           *ebpf.ProgramSpec `ebpf:"kprobe_sys_exit"`
	KprobeTcpConnect        *ebpf.ProgramSpec `ebpf:"kprobe_tcp_connect"`
	KprobeTcpRcvEstablished *ebpf.ProgramSpec `ebpf:"kprobe_tcp_rcv_established"`
	KprobeTcpRecvmsg        *ebpf.ProgramSpec `ebpf:"kprobe_tcp_recvmsg"`
	KprobeTcpSendmsg        *ebpf.ProgramSpec `ebpf:"kprobe_tcp_sendmsg"`
	KretprobeSockAlloc      *ebpf.ProgramSpec `ebpf:"kretprobe_sock_alloc"`
	KretprobeSysAccept4     *ebpf.ProgramSpec `ebpf:"kretprobe_sys_accept4"`
	KretprobeSysClone       *ebpf.ProgramSpec `ebpf:"kretprobe_sys_clone"`
	KretprobeSysConnect     *ebpf.ProgramSpec `ebpf:"kretprobe_sys_connect"`
	KretprobeTcpRecvmsg     *ebpf.ProgramSpec `ebpf:"kretprobe_tcp_recvmsg"`
	SocketHttpFilter        *ebpf.ProgramSpec `ebpf:"socket__http_filter"`
}

// bpf_tpMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tpMapSpecs struct {
	ActiveAcceptArgs    *ebpf.MapSpec `ebpf:"active_accept_args"`
	ActiveConnectArgs   *ebpf.MapSpec `ebpf:"active_connect_args"`
	ActiveRecvArgs      *ebpf.MapSpec `ebpf:"active_recv_args"`
	ActiveSslHandshakes *ebpf.MapSpec `ebpf:"active_ssl_handshakes"`
	ActiveSslReadArgs   *ebpf.MapSpec `ebpf:"active_ssl_read_args"`
	ActiveSslWriteArgs  *ebpf.MapSpec `ebpf:"active_ssl_write_args"`
	CloneMap            *ebpf.MapSpec `ebpf:"clone_map"`
	Events              *ebpf.MapSpec `ebpf:"events"`
	FilteredConnections *ebpf.MapSpec `ebpf:"filtered_connections"`
	HttpInfoMem         *ebpf.MapSpec `ebpf:"http_info_mem"`
	OngoingHttp         *ebpf.MapSpec `ebpf:"ongoing_http"`
	OngoingHttpFallback *ebpf.MapSpec `ebpf:"ongoing_http_fallback"`
	PidCache            *ebpf.MapSpec `ebpf:"pid_cache"`
	PidTidToConn        *ebpf.MapSpec `ebpf:"pid_tid_to_conn"`
	ServerTraces        *ebpf.MapSpec `ebpf:"server_traces"`
	SslToConn           *ebpf.MapSpec `ebpf:"ssl_to_conn"`
	SslToPidTid         *ebpf.MapSpec `ebpf:"ssl_to_pid_tid"`
	TpCharBufMem        *ebpf.MapSpec `ebpf:"tp_char_buf_mem"`
	TpInfoMem           *ebpf.MapSpec `ebpf:"tp_info_mem"`
	TraceMap            *ebpf.MapSpec `ebpf:"trace_map"`
	ValidPids           *ebpf.MapSpec `ebpf:"valid_pids"`
}

// bpf_tpObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tpObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tpObjects struct {
	bpf_tpPrograms
	bpf_tpMaps
}

func (o *bpf_tpObjects) Close() error {
	return _Bpf_tpClose(
		&o.bpf_tpPrograms,
		&o.bpf_tpMaps,
	)
}

// bpf_tpMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tpObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tpMaps struct {
	ActiveAcceptArgs    *ebpf.Map `ebpf:"active_accept_args"`
	ActiveConnectArgs   *ebpf.Map `ebpf:"active_connect_args"`
	ActiveRecvArgs      *ebpf.Map `ebpf:"active_recv_args"`
	ActiveSslHandshakes *ebpf.Map `ebpf:"active_ssl_handshakes"`
	ActiveSslReadArgs   *ebpf.Map `ebpf:"active_ssl_read_args"`
	ActiveSslWriteArgs  *ebpf.Map `ebpf:"active_ssl_write_args"`
	CloneMap            *ebpf.Map `ebpf:"clone_map"`
	Events              *ebpf.Map `ebpf:"events"`
	FilteredConnections *ebpf.Map `ebpf:"filtered_connections"`
	HttpInfoMem         *ebpf.Map `ebpf:"http_info_mem"`
	OngoingHttp         *ebpf.Map `ebpf:"ongoing_http"`
	OngoingHttpFallback *ebpf.Map `ebpf:"ongoing_http_fallback"`
	PidCache            *ebpf.Map `ebpf:"pid_cache"`
	PidTidToConn        *ebpf.Map `ebpf:"pid_tid_to_conn"`
	ServerTraces        *ebpf.Map `ebpf:"server_traces"`
	SslToConn           *ebpf.Map `ebpf:"ssl_to_conn"`
	SslToPidTid         *ebpf.Map `ebpf:"ssl_to_pid_tid"`
	TpCharBufMem        *ebpf.Map `ebpf:"tp_char_buf_mem"`
	TpInfoMem           *ebpf.Map `ebpf:"tp_info_mem"`
	TraceMap            *ebpf.Map `ebpf:"trace_map"`
	ValidPids           *ebpf.Map `ebpf:"valid_pids"`
}

func (m *bpf_tpMaps) Close() error {
	return _Bpf_tpClose(
		m.ActiveAcceptArgs,
		m.ActiveConnectArgs,
		m.ActiveRecvArgs,
		m.ActiveSslHandshakes,
		m.ActiveSslReadArgs,
		m.ActiveSslWriteArgs,
		m.CloneMap,
		m.Events,
		m.FilteredConnections,
		m.HttpInfoMem,
		m.OngoingHttp,
		m.OngoingHttpFallback,
		m.PidCache,
		m.PidTidToConn,
		m.ServerTraces,
		m.SslToConn,
		m.SslToPidTid,
		m.TpCharBufMem,
		m.TpInfoMem,
		m.TraceMap,
		m.ValidPids,
	)
}

// bpf_tpPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tpObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tpPrograms struct {
	KprobeSysExit           *ebpf.Program `ebpf:"kprobe_sys_exit"`
	KprobeTcpConnect        *ebpf.Program `ebpf:"kprobe_tcp_connect"`
	KprobeTcpRcvEstablished *ebpf.Program `ebpf:"kprobe_tcp_rcv_established"`
	KprobeTcpRecvmsg        *ebpf.Program `ebpf:"kprobe_tcp_recvmsg"`
	KprobeTcpSendmsg        *ebpf.Program `ebpf:"kprobe_tcp_sendmsg"`
	KretprobeSockAlloc      *ebpf.Program `ebpf:"kretprobe_sock_alloc"`
	KretprobeSysAccept4     *ebpf.Program `ebpf:"kretprobe_sys_accept4"`
	KretprobeSysClone       *ebpf.Program `ebpf:"kretprobe_sys_clone"`
	KretprobeSysConnect     *ebpf.Program `ebpf:"kretprobe_sys_connect"`
	KretprobeTcpRecvmsg     *ebpf.Program `ebpf:"kretprobe_tcp_recvmsg"`
	SocketHttpFilter        *ebpf.Program `ebpf:"socket__http_filter"`
}

func (p *bpf_tpPrograms) Close() error {
	return _Bpf_tpClose(
		p.KprobeSysExit,
		p.KprobeTcpConnect,
		p.KprobeTcpRcvEstablished,
		p.KprobeTcpRecvmsg,
		p.KprobeTcpSendmsg,
		p.KretprobeSockAlloc,
		p.KretprobeSysAccept4,
		p.KretprobeSysClone,
		p.KretprobeSysConnect,
		p.KretprobeTcpRecvmsg,
		p.SocketHttpFilter,
	)
}

func _Bpf_tpClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_tp_bpfel_x86.o
var _Bpf_tpBytes []byte
