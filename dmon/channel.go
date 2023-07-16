package dmon

type Channel string

const (
	NetworkChannel   Channel = "dmon_network_out"
	StructureChannel Channel = "dmon_structure_out"
	MergedChannel    Channel = "dmon_merged_out"
)
