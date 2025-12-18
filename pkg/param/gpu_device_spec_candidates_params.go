// Copyright (c) ZStack.io, Inc.

package param

// GetGpuDeviceSpecCandidatesDetailParam GetGpuDeviceSpecCandidates详细参数
type GetGpuDeviceSpecCandidatesDetailParam struct {
	rest []string `json:"clusterUuids,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest []string `json:"vmInstanceUuids,omitempty"`
}

// GetGpuDeviceSpecCandidatesParam GetGpuDeviceSpecCandidates请求参数
type GetGpuDeviceSpecCandidatesParam struct {
	BaseParam
	Params GetGpuDeviceSpecCandidatesDetailParam `json:"params"` // 详细参数
}

