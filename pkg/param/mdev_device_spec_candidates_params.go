// Copyright (c) ZStack.io, Inc.

package param

// GetMdevDeviceSpecCandidatesDetailParam GetMdevDeviceSpecCandidates详细参数
type GetMdevDeviceSpecCandidatesDetailParam struct {
	rest []string `json:"clusterUuids,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest []string `json:"vmInstanceUuids,omitempty"`
	rest []string `json:"types,omitempty"`
}

// GetMdevDeviceSpecCandidatesParam GetMdevDeviceSpecCandidates请求参数
type GetMdevDeviceSpecCandidatesParam struct {
	BaseParam
	Params GetMdevDeviceSpecCandidatesDetailParam `json:"params"` // 详细参数
}

