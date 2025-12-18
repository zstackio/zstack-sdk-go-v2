// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateZonesClustersHostsForCreatingVmDetailParam GetCandidateZonesClustersHostsForCreatingVm详细参数
type GetCandidateZonesClustersHostsForCreatingVmDetailParam struct {
	rest string `json:"instanceOfferingUuid,omitempty"`
	rest string `json:"imageUuid" validate:"required"` // 必填
	rest []string `json:"l3NetworkUuids" validate:"required"` // 必填
	rest string `json:"rootDiskOfferingUuid,omitempty"`
	rest int64 `json:"rootDiskSize,omitempty"`
	rest []string `json:"dataDiskOfferingUuids,omitempty"`
	rest int `json:"cpuNum,omitempty"`
	rest int64 `json:"memorySize,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"defaultL3NetworkUuid,omitempty"`
}

// GetCandidateZonesClustersHostsForCreatingVmParam GetCandidateZonesClustersHostsForCreatingVm请求参数
type GetCandidateZonesClustersHostsForCreatingVmParam struct {
	BaseParam
	Params GetCandidateZonesClustersHostsForCreatingVmDetailParam `json:"params"` // 详细参数
}

