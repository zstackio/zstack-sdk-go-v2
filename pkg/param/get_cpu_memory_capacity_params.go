// Copyright (c) ZStack.io, Inc.

package param

// GetCpuMemoryCapacityDetailParam GetCpuMemoryCapacity detail param
type GetCpuMemoryCapacityDetailParam struct {
	ZoneUuids []string `json:"zoneUuids,omitempty"`
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	HostUuids []string `json:"hostUuids,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	All bool `json:"all,omitempty"`
}

// GetCpuMemoryCapacityParam GetCpuMemoryCapacity request param
type GetCpuMemoryCapacityParam struct {
	BaseParam
	Params GetCpuMemoryCapacityDetailParam `json:"params"`
}
