// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package nethttp

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

type bpf_tpFramerFuncInvocationT struct {
	FramerPtr uint64
	Tp        bpf_tpTpInfoT
}

type bpf_tpGoroutineMetadata struct {
	Parent    uint64
	Timestamp uint64
}

type bpf_tpHttpConnectionMetadataT struct {
	Pid struct {
		HostPid   uint32
		UserPid   uint32
		Namespace uint32
	}
	Type uint8
}

type bpf_tpHttpFuncInvocationT struct {
	StartMonotimeNs uint64
	ReqPtr          uint64
	Tp              bpf_tpTpInfoT
}

type bpf_tpPidConnectionInfoT struct {
	Conn bpf_tpConnectionInfoT
	Pid  uint32
}

type bpf_tpPidKeyT struct {
	Pid       uint32
	Namespace uint32
}

type bpf_tpTpInfoPidT struct {
	Tp    bpf_tpTpInfoT
	Pid   uint32
	Valid uint8
	_     [3]byte
}

type bpf_tpTpInfoT struct {
	TraceId  [16]uint8
	SpanId   [8]uint8
	ParentId [8]uint8
	Ts       uint64
	Flags    uint8
	_        [7]byte
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
	UprobeServeHTTP                           *ebpf.ProgramSpec `ebpf:"uprobe_ServeHTTP"`
	UprobeWriteHeader                         *ebpf.ProgramSpec `ebpf:"uprobe_WriteHeader"`
	UprobeConnServe                           *ebpf.ProgramSpec `ebpf:"uprobe_connServe"`
	UprobeConnServeRet                        *ebpf.ProgramSpec `ebpf:"uprobe_connServeRet"`
	UprobeHttp2FramerWriteHeaders             *ebpf.ProgramSpec `ebpf:"uprobe_http2FramerWriteHeaders"`
	UprobeHttp2FramerWriteHeadersReturns      *ebpf.ProgramSpec `ebpf:"uprobe_http2FramerWriteHeaders_returns"`
	UprobeHttp2ResponseWriterStateWriteHeader *ebpf.ProgramSpec `ebpf:"uprobe_http2ResponseWriterStateWriteHeader"`
	UprobeHttp2RoundTrip                      *ebpf.ProgramSpec `ebpf:"uprobe_http2RoundTrip"`
	UprobePersistConnRoundTrip                *ebpf.ProgramSpec `ebpf:"uprobe_persistConnRoundTrip"`
	UprobeReadRequestReturns                  *ebpf.ProgramSpec `ebpf:"uprobe_readRequestReturns"`
	UprobeRoundTrip                           *ebpf.ProgramSpec `ebpf:"uprobe_roundTrip"`
	UprobeRoundTripReturn                     *ebpf.ProgramSpec `ebpf:"uprobe_roundTripReturn"`
	UprobeWriteSubset                         *ebpf.ProgramSpec `ebpf:"uprobe_writeSubset"`
}

// bpf_tpMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tpMapSpecs struct {
	Events                       *ebpf.MapSpec `ebpf:"events"`
	FilteredConnections          *ebpf.MapSpec `ebpf:"filtered_connections"`
	FramerInvocationMap          *ebpf.MapSpec `ebpf:"framer_invocation_map"`
	GoTraceMap                   *ebpf.MapSpec `ebpf:"go_trace_map"`
	GolangMapbucketStorageMap    *ebpf.MapSpec `ebpf:"golang_mapbucket_storage_map"`
	HeaderReqMap                 *ebpf.MapSpec `ebpf:"header_req_map"`
	Http2ReqMap                  *ebpf.MapSpec `ebpf:"http2_req_map"`
	OngoingGoroutines            *ebpf.MapSpec `ebpf:"ongoing_goroutines"`
	OngoingHttpClientRequests    *ebpf.MapSpec `ebpf:"ongoing_http_client_requests"`
	OngoingHttpServerConnections *ebpf.MapSpec `ebpf:"ongoing_http_server_connections"`
	OngoingHttpServerRequests    *ebpf.MapSpec `ebpf:"ongoing_http_server_requests"`
	PidCache                     *ebpf.MapSpec `ebpf:"pid_cache"`
	TraceMap                     *ebpf.MapSpec `ebpf:"trace_map"`
	ValidPids                    *ebpf.MapSpec `ebpf:"valid_pids"`
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
	Events                       *ebpf.Map `ebpf:"events"`
	FilteredConnections          *ebpf.Map `ebpf:"filtered_connections"`
	FramerInvocationMap          *ebpf.Map `ebpf:"framer_invocation_map"`
	GoTraceMap                   *ebpf.Map `ebpf:"go_trace_map"`
	GolangMapbucketStorageMap    *ebpf.Map `ebpf:"golang_mapbucket_storage_map"`
	HeaderReqMap                 *ebpf.Map `ebpf:"header_req_map"`
	Http2ReqMap                  *ebpf.Map `ebpf:"http2_req_map"`
	OngoingGoroutines            *ebpf.Map `ebpf:"ongoing_goroutines"`
	OngoingHttpClientRequests    *ebpf.Map `ebpf:"ongoing_http_client_requests"`
	OngoingHttpServerConnections *ebpf.Map `ebpf:"ongoing_http_server_connections"`
	OngoingHttpServerRequests    *ebpf.Map `ebpf:"ongoing_http_server_requests"`
	PidCache                     *ebpf.Map `ebpf:"pid_cache"`
	TraceMap                     *ebpf.Map `ebpf:"trace_map"`
	ValidPids                    *ebpf.Map `ebpf:"valid_pids"`
}

func (m *bpf_tpMaps) Close() error {
	return _Bpf_tpClose(
		m.Events,
		m.FilteredConnections,
		m.FramerInvocationMap,
		m.GoTraceMap,
		m.GolangMapbucketStorageMap,
		m.HeaderReqMap,
		m.Http2ReqMap,
		m.OngoingGoroutines,
		m.OngoingHttpClientRequests,
		m.OngoingHttpServerConnections,
		m.OngoingHttpServerRequests,
		m.PidCache,
		m.TraceMap,
		m.ValidPids,
	)
}

// bpf_tpPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tpObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tpPrograms struct {
	UprobeServeHTTP                           *ebpf.Program `ebpf:"uprobe_ServeHTTP"`
	UprobeWriteHeader                         *ebpf.Program `ebpf:"uprobe_WriteHeader"`
	UprobeConnServe                           *ebpf.Program `ebpf:"uprobe_connServe"`
	UprobeConnServeRet                        *ebpf.Program `ebpf:"uprobe_connServeRet"`
	UprobeHttp2FramerWriteHeaders             *ebpf.Program `ebpf:"uprobe_http2FramerWriteHeaders"`
	UprobeHttp2FramerWriteHeadersReturns      *ebpf.Program `ebpf:"uprobe_http2FramerWriteHeaders_returns"`
	UprobeHttp2ResponseWriterStateWriteHeader *ebpf.Program `ebpf:"uprobe_http2ResponseWriterStateWriteHeader"`
	UprobeHttp2RoundTrip                      *ebpf.Program `ebpf:"uprobe_http2RoundTrip"`
	UprobePersistConnRoundTrip                *ebpf.Program `ebpf:"uprobe_persistConnRoundTrip"`
	UprobeReadRequestReturns                  *ebpf.Program `ebpf:"uprobe_readRequestReturns"`
	UprobeRoundTrip                           *ebpf.Program `ebpf:"uprobe_roundTrip"`
	UprobeRoundTripReturn                     *ebpf.Program `ebpf:"uprobe_roundTripReturn"`
	UprobeWriteSubset                         *ebpf.Program `ebpf:"uprobe_writeSubset"`
}

func (p *bpf_tpPrograms) Close() error {
	return _Bpf_tpClose(
		p.UprobeServeHTTP,
		p.UprobeWriteHeader,
		p.UprobeConnServe,
		p.UprobeConnServeRet,
		p.UprobeHttp2FramerWriteHeaders,
		p.UprobeHttp2FramerWriteHeadersReturns,
		p.UprobeHttp2ResponseWriterStateWriteHeader,
		p.UprobeHttp2RoundTrip,
		p.UprobePersistConnRoundTrip,
		p.UprobeReadRequestReturns,
		p.UprobeRoundTrip,
		p.UprobeRoundTripReturn,
		p.UprobeWriteSubset,
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
//go:embed bpf_tp_bpfel_arm64.o
var _Bpf_tpBytes []byte
