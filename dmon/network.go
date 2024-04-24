package dmon

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/v1dmon/payload/common"
)

// NOTE: FT_RELATIVE_TIME is a Wireshark/libpcap type for time delta values
// Precision: example value: 1.123456789 => seconds.nanoseconds
// Converted by default to float64
type NetworkGeneral struct {
	common.Header
	IPSrc             string  `json:"IPSrc"`             // Source Address, ip.src
	IPDst             string  `json:"IPDst"`             // Destination Address, ip.dst
	Protocol          string  `json:"Protocol"`          // Protocol, _ws.col.protocol
	FrameNumber       uint64  `json:"FrameNumber"`       // Frame number, frame.number (uint32)
	FrameTimeDelta    float64 `json:"FrameTimeDelta"`    // Time delta from previous captured frame, frame.time_delta (FT_RELATIVE_TIME->float64)
	FrameTimeRelative float64 `json:"FrameTimeRelative"` // Time since reference or first frame, frame.time_relative (FT_RELATIVE_TIME->float64)
	TCPStream         uint64  `json:"TCPStream"`         // Stream index, tcp.stream (uint32)
	TCPSrcPort        string  `json:"TCPSrcPort"`        // Source Port, tcp.srcport
	TCPDstPort        string  `json:"TCPDstPort"`        // Destination Port, tcp.dstport
	TCPTimeDelta      float64 `json:"TCPTimeDelta"`      // Time since previous frame in this TCP stream, tcp.time_delta (FT_RELATIVE_TIME->float64)
	TCPTimeRelative   float64 `json:"TCPTimRelative"`    // Time since first frame in this TCP stream, tcp.time_relative (FT_RELATIVE_TIME->float64)
	IsRequest         string  `json:"IsRequest"`         // Request, http.request (bool)
	RequestNumber     string  `json:"RequestNumber"`     // Request number, http.request_number (uint32->string, WARN: "" == NaN)
	RequestMethod     string  `json:"RequestMethod"`     // Request method, http.request.method
	RequestFullURI    string  `json:"RequestFullURI"`    // Full request URI, http.request.full_uri
	IsResponse        string  `json:"IsResponse"`        // Response, http.response (bool)
	ResponseNumber    string  `json:"ResponseNumber"`    // Response number, http.response_number (uint32)
	ResponseCode      string  `json:"ResponseCode"`      // Status code, http.response.code (uint24bit->string, WARN: "" == NaN)
	ResponseCodeDesc  string  `json:"ResponseCodeDesc"`  // Status code description, http.response.code.desc
	ResponseForURI    string  `json:"ResponseForURI"`    // Request URI, http.response_for.uri
	ResponseTime      string  `json:"ResponseTime"`      // Time since request, http.time (FT_RELATIVE_TIME->string, WARN: "" == NaN)
}

func NewNetworkGeneral(packet []string) (*NetworkGeneral, error) {
	networkGeneral := NetworkGeneral{}
	t, err := time.Parse("2006-01-02 15:04:05.000000000", packet[0])
	if err != nil {
		return nil, err
	}
	networkGeneral.Timestamp = t.Format(time.RFC3339)
	networkGeneral.Type = string(NetworkType)
	networkGeneral.SubType = string(NetworkGeneralSubType)

	networkGeneral.IPSrc = packet[1]
	networkGeneral.IPDst = packet[2]
	networkGeneral.Protocol = packet[3]

	// frame
	u64, err := strconv.ParseUint(packet[4], 10, 64)
	if err != nil {
		return nil, err
	}
	networkGeneral.FrameNumber = u64
	f64, err := strconv.ParseFloat(packet[5], 64)
	if err != nil {
		return nil, err
	}
	networkGeneral.FrameTimeDelta = f64
	f64, err = strconv.ParseFloat(packet[6], 64)
	if err != nil {
		return nil, err
	}
	networkGeneral.FrameTimeRelative = f64

	// TCP
	u64, err = strconv.ParseUint(packet[4], 10, 64)
	if err != nil {
		return nil, err
	}
	networkGeneral.TCPStream = u64
	networkGeneral.TCPSrcPort = packet[8]
	networkGeneral.TCPDstPort = packet[9]
	f64, err = strconv.ParseFloat(packet[10], 64)
	if err != nil {
		return nil, err
	}
	networkGeneral.TCPTimeDelta = f64
	f64, err = strconv.ParseFloat(packet[11], 64)
	if err != nil {
		return nil, err
	}
	networkGeneral.TCPTimeRelative = f64

	// HTTP request
	networkGeneral.IsRequest = packet[12]
	networkGeneral.RequestNumber = packet[13]
	networkGeneral.RequestMethod = packet[14]
	networkGeneral.RequestFullURI = packet[15]

	// HTTP response
	networkGeneral.IsResponse = packet[16]
	networkGeneral.ResponseNumber = packet[17]
	networkGeneral.ResponseCode = packet[18]
	networkGeneral.ResponseCodeDesc = packet[19]
	networkGeneral.ResponseForURI = packet[20]
	networkGeneral.ResponseTime = packet[21]

	return &networkGeneral, nil
}

func (n *NetworkGeneral) Marshal() ([]byte, error) {
	enc, err := json.Marshal(n)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func (n *NetworkGeneral) Display(e func() *zerolog.Event) {
	e().
		Str("Type", n.Type).
		Str("SubType", n.SubType).
		Str("IPSrc", n.IPSrc).
		Str("IPDst", n.IPDst).
		Str("Protocol", n.Protocol).
		Uint64("FrameNumber", n.FrameNumber).
		Float64("FrameTimeDelta", n.FrameTimeDelta).
		Float64("FrameTimeRelative", n.FrameTimeRelative).
		Uint64("TCPStream", n.TCPStream).
		Str("TCPSrcPort", n.TCPSrcPort).
		Str("TCPDstPort", n.TCPDstPort).
		Float64("TCPTimeDelta", n.TCPTimeDelta).
		Float64("TCPTimeRelative", n.TCPTimeRelative).
		Str("IsRequest", n.IsRequest).
		Str("RequestNumber", n.RequestNumber).
		Str("RequestMethod", n.RequestMethod).
		Str("RequestFullURI", n.RequestFullURI).
		Str("IsResponse", n.IsResponse).
		Str("ResponseNumber", n.ResponseNumber).
		Str("ResponseCode", n.ResponseCode).
		Str("ResponseCodeDesc", n.ResponseCodeDesc).
		Str("ResponseForURI", n.ResponseForURI).
		Str("ResponseTime", n.ResponseTime).
		Send()
}
