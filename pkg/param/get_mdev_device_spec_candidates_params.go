// Copyright (c) ZStack.io, Inc.

package param

// GetMdevDeviceSpecCandidatesDetailParam GetMdevDeviceSpecCandidates detail param
type GetMdevDeviceSpecCandidatesDetailParam struct {
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	VmInstanceUuids []string `json:"vmInstanceUuids,omitempty"`
	Types []string `json:"types,omitempty"`
}

// GetMdevDeviceSpecCandidatesParam GetMdevDeviceSpecCandidates request param
type GetMdevDeviceSpecCandidatesParam struct {
	BaseParam
	Params GetMdevDeviceSpecCandidatesDetailParam `json:"params"`
}
