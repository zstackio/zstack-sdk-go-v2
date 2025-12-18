// Copyright (c) ZStack.io, Inc.

package param

// GetPciDeviceCandidatesForAttachingVmDetailParam GetPciDeviceCandidatesForAttachingVm详细参数
type GetPciDeviceCandidatesForAttachingVmDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest []string `json:"types,omitempty"`
	rest []string `json:"pciSpecUuids,omitempty"`
}

// GetPciDeviceCandidatesForAttachingVmParam GetPciDeviceCandidatesForAttachingVm请求参数
type GetPciDeviceCandidatesForAttachingVmParam struct {
	BaseParam
	Params GetPciDeviceCandidatesForAttachingVmDetailParam `json:"params"` // 详细参数
}

