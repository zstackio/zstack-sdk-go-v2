// Copyright (c) ZStack.io, Inc.

package param

// GetPciDeviceSpecCandidatesDetailParam GetPciDeviceSpecCandidates详细参数
type GetPciDeviceSpecCandidatesDetailParam struct {
	rest []string `json:"clusterUuids,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest []string `json:"vmInstanceUuids,omitempty"`
	rest []string `json:"types,omitempty"`
}

// GetPciDeviceSpecCandidatesParam GetPciDeviceSpecCandidates请求参数
type GetPciDeviceSpecCandidatesParam struct {
	BaseParam
	Params GetPciDeviceSpecCandidatesDetailParam `json:"params"` // 详细参数
}

