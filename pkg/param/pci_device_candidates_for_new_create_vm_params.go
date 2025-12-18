// Copyright (c) ZStack.io, Inc.

package param

// GetPciDeviceCandidatesForNewCreateVmDetailParam GetPciDeviceCandidatesForNewCreateVm详细参数
type GetPciDeviceCandidatesForNewCreateVmDetailParam struct {
	rest string `json:"hostUuid,omitempty"`
	rest []string `json:"clusterUuids,omitempty"`
	rest []string `json:"types,omitempty"`
}

// GetPciDeviceCandidatesForNewCreateVmParam GetPciDeviceCandidatesForNewCreateVm请求参数
type GetPciDeviceCandidatesForNewCreateVmParam struct {
	BaseParam
	Params GetPciDeviceCandidatesForNewCreateVmDetailParam `json:"params"` // 详细参数
}

