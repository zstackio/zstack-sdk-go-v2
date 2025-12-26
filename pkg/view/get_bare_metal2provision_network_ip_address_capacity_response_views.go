// Copyright (c) ZStack.io, Inc.

package view

// GetBareMetal2ProvisionNetworkIpAddressCapacityView GetBareMetal2ProvisionNetworkIpAddressCapacity
type GetBareMetal2ProvisionNetworkIpAddressCapacityView struct {
	CapacityData []BareMetal2ProvisionNetworkIpCapacityView `json:"capacityData,omitempty"`
	Success bool `json:"success,omitempty"`
}

