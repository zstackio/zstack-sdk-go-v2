// Copyright (c) ZStack.io, Inc.

package param

// GetPciDeviceSpecCandidatesDetailParam GetPciDeviceSpecCandidates detail param
type GetPciDeviceSpecCandidatesDetailParam struct {
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	VmInstanceUuids []string `json:"vmInstanceUuids,omitempty"`
	Types []string `json:"types,omitempty"`
}

// GetPciDeviceSpecCandidatesParam GetPciDeviceSpecCandidates request param
type GetPciDeviceSpecCandidatesParam struct {
	BaseParam
	Params GetPciDeviceSpecCandidatesDetailParam `json:"params"`
}
