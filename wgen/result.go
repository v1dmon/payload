package wgen

import (
	"net/http"
	"time"

	zl "github.com/rs/zerolog/log"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

type HTTPResult struct {
	Header
	Attack   string        `json:"attack"`    // attack string
	Seq      uint64        `json:"seq"`       // packet sequence number
	Code     uint16        `json:"code"`      // http code
	Latency  time.Duration `json:"latency"`   // latency time delta
	BytesIn  uint64        `json:"bytes_in"`  // rx bytes
	BytesOut uint64        `json:"bytes_out"` // tx bytes
	Error    string        `json:"error"`     // error reason
	Body     []byte        `json:"body"`      // http body
	Method   string        `json:"method"`    // http method
	URL      string        `json:"url"`       // target url
	Headers  http.Header   `json:"headers"`   // http headers
}

func NewHTTPResult(result *vegeta.Result) (*HTTPResult, error) {
	httpResult := HTTPResult{}
	httpResult.Timestamp = result.Timestamp.UTC().Format(time.RFC3339)
	httpResult.Type = string(HttpType)
	httpResult.SubType = string(HttpResultSubType)
	httpResult.Attack = result.Attack
	httpResult.Seq = result.Seq
	httpResult.Code = result.Code
	httpResult.Latency = result.Latency
	httpResult.BytesIn = result.BytesIn
	httpResult.BytesOut = result.BytesOut
	httpResult.Error = result.Error
	httpResult.Body = result.Body
	httpResult.Method = result.Method
	httpResult.URL = result.URL
	httpResult.Headers = result.Headers
	return &httpResult, nil
}

func (n *HTTPResult) Marshal() ([]byte, error) {
	return marshal(n)
}

func (n *HTTPResult) Display() {
	zl.Info().
		Str("Timestamp", n.Timestamp).
		Str("Type", n.Type).
		Str("SubType", n.SubType).
		Str("Attack", n.Attack).
		Uint64("Seq", n.Seq).
		Uint16("Code", n.Code).
		Str("Latency", n.Latency.String()).
		Uint64("BytesOut", n.BytesOut).
		Uint64("BytesIn", n.BytesIn).
		Str("Error", n.Error).
		Str("Method", n.Method).
		Str("URL", n.URL).
		Send()
}
