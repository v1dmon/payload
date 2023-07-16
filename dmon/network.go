package dmon

import (
	"encoding/json"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/v1dmon/payload/common"
)

type NetworkGeneral struct {
	common.Header
	SendIP    string  `json:"SendIP"`
	RecvIP    string  `json:"RecvIP"`
	Protocol  string  `json:"Protocol"`
	TimeDelta float64 `json:"TimeDelta"`
}

func NewNetworkGeneral(packet []string) (*NetworkGeneral, error) {
	networkGeneral := NetworkGeneral{}
	networkGeneral.Timestamp = packet[0] + " " + packet[1]
	networkGeneral.Type = string(NetworkType)
	networkGeneral.SubType = string(NetworkGeneralSubType)
	networkGeneral.SendIP = packet[2]
	networkGeneral.RecvIP = packet[3]
	networkGeneral.Protocol = packet[4]
	timeDelta, err := strconv.ParseFloat(packet[len(packet)-1], 64)
	if err != nil {
		return nil, err
	}
	networkGeneral.TimeDelta = timeDelta
	return &networkGeneral, nil
}

func (n *NetworkGeneral) Marshal() ([]byte, error) {
	enc, err := json.Marshal(n)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func (n *NetworkGeneral) Display(l *zerolog.Logger) {
	l.Info().
		Str("Timestamp", n.Timestamp).
		Str("Type", n.Type).
		Str("SubType", n.SubType).
		Str("SendIP", n.SendIP).
		Str("RecvIP", n.RecvIP).
		Str("Protocol", n.Protocol).
		Float64("TimeDelta", n.TimeDelta).
		Send()
}
