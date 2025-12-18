// Copyright (c) ZStack.io, Inc.

package param

// GetCpuMemoryCapacityDetailParam GetCpuMemoryCapacity详细参数
type GetCpuMemoryCapacityDetailParam struct {
	rest []string `json:"zoneUuids,omitempty"`
	rest []string `json:"clusterUuids,omitempty"`
	rest []string `json:"hostUuids,omitempty"`
	rest string `json:"hypervisorType,omitempty"`
	rest bool `json:"all,omitempty"`
}

// GetCpuMemoryCapacityParam GetCpuMemoryCapacity请求参数
type GetCpuMemoryCapacityParam struct {
	BaseParam
	Params GetCpuMemoryCapacityDetailParam `json:"params"` // 详细参数
}

