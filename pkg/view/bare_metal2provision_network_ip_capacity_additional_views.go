// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2ProvisionNetworkIpCapacityView BareMetal2ProvisionNetworkIpCapacity
type BareMetal2ProvisionNetworkIpCapacityView struct {
	NetworkUuid          string `json:"networkUuid,omitempty"`
	TotalCapacity        int64  `json:"totalCapacity,omitempty"`
	AvailableCapacity    int64  `json:"availableCapacity,omitempty"`
	GatewayUsedIpNumber  int64  `json:"gatewayUsedIpNumber,omitempty"`
	InstanceUsedIpNumber int64  `json:"instanceUsedIpNumber,omitempty"`
}
