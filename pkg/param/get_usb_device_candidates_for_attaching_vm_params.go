// Copyright (c) ZStack.io, Inc.

package param

// GetUsbDeviceCandidatesForAttachingVmDetailParam GetUsbDeviceCandidatesForAttachingVm detail param
type GetUsbDeviceCandidatesForAttachingVmDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	AttachType string `json:"attachType,omitempty"`
}

// GetUsbDeviceCandidatesForAttachingVmParam GetUsbDeviceCandidatesForAttachingVm request param
type GetUsbDeviceCandidatesForAttachingVmParam struct {
	BaseParam
	Params GetUsbDeviceCandidatesForAttachingVmDetailParam `json:"params"`
}
