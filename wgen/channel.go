package wgen

type Channel string

const (
	ResultsChannel Channel = "wgen_results_out"
	MetricsChannel Channel = "wgen_metrics_out"
	MergedChannel  Channel = "wgen_merged_out"
)
