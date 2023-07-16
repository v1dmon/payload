package dmon

import (
	"github.com/v1dmon/payload/common"
)

const (
	NetworkType   common.PayloadType = "network"
	StructureType common.PayloadType = "structure"
)

const (
	NetworkGeneralSubType     common.PayloadSubType = "general"
	StructureHostSubType      common.PayloadSubType = "host"
	StructureNetworkSubType   common.PayloadSubType = "network"
	StructureContainerSubType common.PayloadSubType = "container"
)
