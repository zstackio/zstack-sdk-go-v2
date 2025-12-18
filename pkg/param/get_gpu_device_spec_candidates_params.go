// Copyright (c) ZStack.io, Inc.

package param

// GetGpuDeviceSpecCandidatesDetailParam GetGpuDeviceSpecCandidates detail param
type GetGpuDeviceSpecCandidatesDetailParam struct {
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	VmInstanceUuids []string `json:"vmInstanceUuids,omitempty"`
}

// GetGpuDeviceSpecCandidatesParam GetGpuDeviceSpecCandidates request param
type GetGpuDeviceSpecCandidatesParam struct {
	BaseParam
	Params GetGpuDeviceSpecCandidatesDetailParam `json:"params"`
}
