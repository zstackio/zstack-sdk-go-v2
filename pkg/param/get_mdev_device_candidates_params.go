// Copyright (c) ZStack.io, Inc.

package param

// GetMdevDeviceCandidatesDetailParam GetMdevDeviceCandidates detail param
type GetMdevDeviceCandidatesDetailParam struct {
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	Types []string `json:"types,omitempty"`
}

// GetMdevDeviceCandidatesParam GetMdevDeviceCandidates request param
type GetMdevDeviceCandidatesParam struct {
	BaseParam
	Params GetMdevDeviceCandidatesDetailParam `json:"params"`
}
