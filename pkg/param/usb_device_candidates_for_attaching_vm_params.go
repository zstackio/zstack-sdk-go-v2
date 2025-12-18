// Copyright (c) ZStack.io, Inc.

package param

// GetUsbDeviceCandidatesForAttachingVmDetailParam GetUsbDeviceCandidatesForAttachingVm详细参数
type GetUsbDeviceCandidatesForAttachingVmDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"attachType,omitempty"`
}

// GetUsbDeviceCandidatesForAttachingVmParam GetUsbDeviceCandidatesForAttachingVm请求参数
type GetUsbDeviceCandidatesForAttachingVmParam struct {
	BaseParam
	Params GetUsbDeviceCandidatesForAttachingVmDetailParam `json:"params"` // 详细参数
}

