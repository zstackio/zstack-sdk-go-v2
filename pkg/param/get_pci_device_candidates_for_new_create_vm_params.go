// Copyright (c) ZStack.io, Inc.

package param

// GetPciDeviceCandidatesForNewCreateVmDetailParam GetPciDeviceCandidatesForNewCreateVm detail param
type GetPciDeviceCandidatesForNewCreateVmDetailParam struct {
	HostUuid string `json:"hostUuid,omitempty"`
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	Types []string `json:"types,omitempty"`
}

// GetPciDeviceCandidatesForNewCreateVmParam GetPciDeviceCandidatesForNewCreateVm request param
type GetPciDeviceCandidatesForNewCreateVmParam struct {
	BaseParam
	Params GetPciDeviceCandidatesForNewCreateVmDetailParam `json:"params"`
}
