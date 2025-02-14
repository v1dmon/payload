package wgen

import (
	"encoding/json"
	// "net/http"
	"time"

	"github.com/rs/zerolog"
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"github.com/v1dmon/payload/common"
)

type HTTPResult struct {
	common.Header
	Attack   string        `json:"attack"`    // attack string
	Seq      uint64        `json:"seq"`       // packet sequence number
	Code     uint16        `json:"code"`      // http code
	Latency  time.Duration `json:"latency"`   // latency time delta
	BytesIn  uint64        `json:"bytes_in"`  // rx bytes
	BytesOut uint64        `json:"bytes_out"` // tx bytes
	Error    string        `json:"error"`     // error reason
	Method   string        `json:"method"`    // http method
	URL      string        `json:"url"`       // target url
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
	httpResult.Method = result.Method
	httpResult.URL = result.URL
	return &httpResult, nil
}

func (n *HTTPResult) Marshal() ([]byte, error) {
	enc, err := json.Marshal(n)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func (n *HTTPResult) Display(e func() *zerolog.Event) {
	e().
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
