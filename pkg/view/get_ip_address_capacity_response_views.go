// Copyright (c) ZStack.io, Inc.

package view

// GetIpAddressCapacityView GetIpAddressCapacity
type GetIpAddressCapacityView struct {
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	UsedIpAddressNumber int64 `json:"usedIpAddressNumber,omitempty"`
	Ipv4TotalCapacity int64 `json:"ipv4TotalCapacity,omitempty"`
	Ipv4AvailableCapacity int64 `json:"ipv4AvailableCapacity,omitempty"`
	Ipv4UsedIpAddressNumber int64 `json:"ipv4UsedIpAddressNumber,omitempty"`
	Ipv6TotalCapacity int64 `json:"ipv6TotalCapacity,omitempty"`
	Ipv6AvailableCapacity int64 `json:"ipv6AvailableCapacity,omitempty"`
	Ipv6UsedIpAddressNumber int64 `json:"ipv6UsedIpAddressNumber,omitempty"`
	CapacityData []IpCapacityDataView `json:"capacityData,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Success bool `json:"success,omitempty"`
}

