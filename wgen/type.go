package wgen

type PayloadType string

const (
	HttpType PayloadType = "http"
)

type PayloadSubType string

const (
	HttpResultSubType  PayloadSubType = "result"
	HttpMetricsSubType PayloadSubType = "metrics"
)
