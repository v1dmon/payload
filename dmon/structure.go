package dmon

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/rs/zerolog"
	"github.com/v1dmon/payload/common"
	"kythe.io/kythe/go/util/datasize"
)

type StructureHost struct {
	common.Header
	OperatingSystem string `json:"OperatingSystem"`
	OSType          string `json:"OSType"`
	Architecture    string `json:"Architecture"`
	Name            string `json:"Name"`
	NCPU            int    `json:"NCPU"`
	MemTotal        int64  `json:"MemTotal"`
	KernelVersion   string `json:"KernelVersion"`
}

func NewStructureHost(host *types.Info) *StructureHost {
	structureHost := StructureHost{}
	structureHost.Timestamp = time.Now().UTC().Format(time.RFC3339)
	structureHost.Type = string(StructureType)
	structureHost.SubType = string(StructureHostSubType)
	structureHost.OperatingSystem = host.OperatingSystem
	structureHost.OSType = host.OSType
	structureHost.Architecture = host.Architecture
	structureHost.Name = host.Name
	structureHost.NCPU = host.NCPU
	structureHost.MemTotal = (host.MemTotal) / (1000 * 1000)
	structureHost.KernelVersion = host.KernelVersion
	return &structureHost
}

func (s *StructureHost) Marshal() ([]byte, error) {
	enc, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func (s *StructureHost) Display(e func() *zerolog.Event) {
	e().
		Str("Timestamp", s.Timestamp).
		Str("Type", s.Type).
		Str("SubType", s.SubType).
		Str("OS", s.OperatingSystem).
		Str("OSType", s.OSType).
		Str("Arch", s.Architecture).
		Str("Name", s.Name).
		Int("NCPU", s.NCPU).
		Int64("Memory", s.MemTotal).
		Str("KernelVer", s.KernelVersion).
		Send()
}

type StructureNetwork struct {
	common.Header
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

func NewStructureNetwork(network *types.NetworkResource) *StructureNetwork {
	structureNetwork := StructureNetwork{}
	structureNetwork.Timestamp = time.Now().UTC().Format(time.RFC3339)
	structureNetwork.Type = string(StructureType)
	structureNetwork.SubType = string(StructureNetworkSubType)
	structureNetwork.ID = network.ID
	structureNetwork.Name = network.Name
	return &structureNetwork
}

func (s *StructureNetwork) Marshal() ([]byte, error) {
	enc, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func (s *StructureNetwork) Display(e func() *zerolog.Event) {
	e().
		Str("Timestamp", s.Timestamp).
		Str("Type", s.Type).
		Str("SubType", s.SubType).
		Str("ID", s.ID[:16]).
		Str("Name", s.Name).
		Send()
}

type StructureContainer struct {
	common.Header
	ID          string                                 `json:"ID"`
	Name        string                                 `json:"Name"`
	Image       string                                 `json:"Image"`
	Stats       StructureContainerStats                `json:"Stats"`
	Locale      string                                 `json:"Locale"`
	Timezone    string                                 `json:"Timezone"`
	IPAddresses map[string]string                      `json:"IPAddresses"`
	Ports       map[string][]StructureContainerNetwork `json:"Ports"`
}

type StructureContainerNetwork struct {
	PrivatePort uint16 `json:"PrivatePort"`
	PublicPort  uint16 `json:"PublicPort"`
	Type        string `json:"Type"`
}

type StructureContainerStats struct {
	CPURawUsage       float64 `json:"CPURawUsage"`       // cpu usage in bytes
	CPUSystemRawUsage float64 `json:"CPUSystemRawUsage"` // system cpu usage in bytes
	CPUsOnline        float64 `json:"CPUOnline"`         // number of online cpus
	CPUPercUsage      string  `json:"CPUPercUsage"`      // cpu usage percentage
	MemoryRawLimit    float64 `json:"MemoryRawLimit"`    // memory limit in bytes
	MemoryRawUsage    float64 `json:"MemoryRawUsage"`    // memory usage in bytes
	MemoryPercUsage   string  `json:"MemoryPercUsage"`   // memory usage percentage
	NetworkRxRaw      float64 `json:"NetworkRxRaw"`      // network rx bytes
	NetworkTxRaw      float64 `json:"NetworkTxRaw"`      // network tx bytes
	BlockIORawRead    float64 `json:"BlockIORawRead"`    // block i/o read bytes
	BlockIORawWrite   float64 `json:"BlockIORawWrite"`   // block i/o write bytes
	PIDsCurrent       float64 `json:"PIDsCurrent"`       // current number of pids
	PIDsLimit         float64 `json:"PIDsLimit"`         // max number of pids
}

func NewStructureContainer(
	container *types.Container,
	containerJSONBase *types.ContainerJSONBase,
	containerStats *StructureContainerStats,
	locale, timezone string,
) *StructureContainer {
	structureContainer := StructureContainer{}
	structureContainer.Timestamp = time.Now().UTC().Format(time.RFC3339)
	structureContainer.Type = string(StructureType)
	structureContainer.SubType = string(StructureContainerSubType)
	structureContainer.ID = containerJSONBase.ID
	structureContainer.Name = containerJSONBase.Name
	structureContainer.Image = containerJSONBase.Image
	structureContainer.Stats = *containerStats
	structureContainer.IPAddresses = make(map[string]string)
	structureContainer.Ports = make(map[string][]StructureContainerNetwork)
	structureContainer.Locale = locale
	structureContainer.Timezone = timezone
	for network, info := range container.NetworkSettings.Networks {
		structureContainer.IPAddresses[network] = info.IPAddress
	}
	for _, port := range container.Ports {
		structureContainer.Ports[port.IP] = append(
			structureContainer.Ports[port.IP],
			StructureContainerNetwork{
				PrivatePort: port.PrivatePort,
				PublicPort:  port.PublicPort,
				Type:        port.Type,
			},
		)
	}
	return &structureContainer
}

func (s *StructureContainer) Marshal() ([]byte, error) {
	enc, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func (s *StructureContainer) Display(e func() *zerolog.Event) {
	e().
		Str("Timestamp", s.Timestamp).
		Str("Type", s.Type).
		Str("SubType", s.SubType).
		Str("ID", s.ID[:16]).
		Str("Name", s.Name).
		Str("Image", s.Image[:16]).
		Str("Locale", s.Locale).
		Str("Timezone", s.Timezone).
		Str("CpuUsage", s.Stats.CPUPercUsage).
		Str("MemoryUsage", s.Stats.MemoryPercUsage).
		Str("NetworkUsage(rx/tx)", fmt.Sprintf("%s/%s", btos(s.Stats.NetworkRxRaw), btos(s.Stats.NetworkTxRaw))).
		Str("BlockIOUsage(r/w)", fmt.Sprintf("%s/%s", btos(s.Stats.BlockIORawRead), btos(s.Stats.BlockIORawWrite))).
		Str("PIDs(#/limit)", fmt.Sprintf("%.0f/%.0f", s.Stats.PIDsCurrent, s.Stats.PIDsLimit)).
		Send()
}

// datasize bytes to string
func btos(n any) string {
	// WARN skipping errors.
	s, _ := datasize.Parse(fmt.Sprintf("%vB", n))
	return s.String()
}
