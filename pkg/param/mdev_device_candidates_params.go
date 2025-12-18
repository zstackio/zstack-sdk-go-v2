// Copyright (c) ZStack.io, Inc.

package param

// GetMdevDeviceCandidatesDetailParam GetMdevDeviceCandidates详细参数
type GetMdevDeviceCandidatesDetailParam struct {
	rest []string `json:"clusterUuids,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest []string `json:"types,omitempty"`
}

// GetMdevDeviceCandidatesParam GetMdevDeviceCandidates请求参数
type GetMdevDeviceCandidatesParam struct {
	BaseParam
	Params GetMdevDeviceCandidatesDetailParam `json:"params"` // 详细参数
}

