package dmon

type PayloadType string

const (
	NetworkType   PayloadType = "network"
	StructureType PayloadType = "structure"
)

type payloadSubType string

const (
	NetworkGeneralSubType     payloadSubType = "general"
	StructureHostSubType      payloadSubType = "host"
	StructureNetworkSubType   payloadSubType = "network"
	StructureContainerSubType payloadSubType = "container"
)
