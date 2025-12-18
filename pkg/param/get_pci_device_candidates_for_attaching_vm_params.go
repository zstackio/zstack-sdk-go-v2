// Copyright (c) ZStack.io, Inc.

package param

// GetPciDeviceCandidatesForAttachingVmDetailParam GetPciDeviceCandidatesForAttachingVm detail param
type GetPciDeviceCandidatesForAttachingVmDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	Types []string `json:"types,omitempty"`
	PciSpecUuids []string `json:"pciSpecUuids,omitempty"`
}

// GetPciDeviceCandidatesForAttachingVmParam GetPciDeviceCandidatesForAttachingVm request param
type GetPciDeviceCandidatesForAttachingVmParam struct {
	BaseParam
	Params GetPciDeviceCandidatesForAttachingVmDetailParam `json:"params"`
}
