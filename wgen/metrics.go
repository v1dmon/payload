package wgen

import (
	"time"

	zl "github.com/rs/zerolog/log"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

type HTTPMetrics struct {
	Header
	Latencies  vegeta.LatencyMetrics `json:"latencies"`  // request latency metrics
	BytesIn    vegeta.ByteMetrics    `json:"bytes_in"`   // incoming byte metrics
	BytesOut   vegeta.ByteMetrics    `json:"bytes_out"`  // outgoing byte metrics
	Earliest   time.Time             `json:"earliest"`   // earliest timestamp in a result set
	Latest     time.Time             `json:"latest"`     // timestamp in a Result set
	End        time.Time             `json:"end"`        // latest timestamp in a Result set plus its latency
	Duration   time.Duration         `json:"duration"`   // duration of the attack
	Wait       time.Duration         `json:"wait"`       // extra time waiting for responses from targets
	Requests   uint64                `json:"requests"`   // total number of requests executed
	Rate       float64               `json:"rate"`       // rate of sent requests per second
	Throughput float64               `json:"throughput"` // rate of successful requests per second
	Success    float64               `json:"success"`    // percentage of non-error responses
}

func NewHTTPMetrics(metrics *vegeta.Metrics) (*HTTPMetrics, error) {
	httpMetrics := HTTPMetrics{}
	httpMetrics.Timestamp = metrics.End.UTC().Format(time.RFC3339)
	httpMetrics.Type = string(HttpType)
	httpMetrics.SubType = string(HttpMetricsSubType)
	httpMetrics.Latencies = metrics.Latencies
	httpMetrics.BytesIn = metrics.BytesIn
	httpMetrics.BytesOut = metrics.BytesOut
	httpMetrics.Earliest = metrics.Earliest
	httpMetrics.Latest = metrics.Latest
	httpMetrics.End = metrics.End
	httpMetrics.Duration = metrics.Duration
	httpMetrics.Wait = metrics.Wait
	httpMetrics.Requests = metrics.Requests
	httpMetrics.Rate = metrics.Rate
	httpMetrics.Throughput = metrics.Throughput
	httpMetrics.Success = metrics.Success
	return &httpMetrics, nil
}

func (n *HTTPMetrics) Marshal() ([]byte, error) {
	return marshal(n)
}

func (n *HTTPMetrics) Display() {
	zl.Info().
		Str("Timestamp", n.Timestamp).
		Str("Type", n.Type).
		Str("SubType", n.SubType).
		Str("Latencies", n.Latencies.Mean.String()).
		Float64("BytesIn", n.BytesIn.Mean).
		Float64("BytesOut", n.BytesOut.Mean).
		Str("Earliest", n.Earliest.Format(time.RFC3339)).
		Str("Latest", n.Latest.Format(time.RFC3339)).
		Str("End", n.End.Format(time.RFC3339)).
		Str("Duration", n.Duration.String()).
		Str("Wait", n.Wait.String()).
		Uint64("Requests", n.Requests).
		Float64("Rate", n.Rate).
		Float64("Throughput", n.Throughput).
		Float64("Success", n.Success).
		Send()
}
