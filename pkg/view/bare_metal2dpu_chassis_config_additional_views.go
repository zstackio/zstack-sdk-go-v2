// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BareMetal2DpuChassisConfigView BareMetal2DpuChassisConfig
type BareMetal2DpuChassisConfigView struct {
	IpmiAddress string `json:"ipmiAddress,omitempty"`
	IpmiPort int `json:"ipmiPort,omitempty"`
	IpmiUsername string `json:"ipmiUsername,omitempty"`
	IpmiPassword string `json:"ipmiPassword,omitempty"`
}

