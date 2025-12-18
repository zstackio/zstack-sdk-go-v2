// Copyright (c) ZStack.io, Inc.

package param

// GetCandidatePrimaryStoragesForCreatingVmDetailParam GetCandidatePrimaryStoragesForCreatingVm详细参数
type GetCandidatePrimaryStoragesForCreatingVmDetailParam struct {
	rest string `json:"imageUuid" validate:"required"` // 必填
	rest []string `json:"l3NetworkUuids" validate:"required"` // 必填
	rest string `json:"rootDiskOfferingUuid,omitempty"`
	rest int64 `json:"rootDiskSize,omitempty"`
	rest []string `json:"dataDiskOfferingUuids,omitempty"`
	rest []int64 `json:"dataDiskSizes,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"defaultL3NetworkUuid,omitempty"`
	rest string `json:"instanceOfferingUuid,omitempty"`
}

// GetCandidatePrimaryStoragesForCreatingVmParam GetCandidatePrimaryStoragesForCreatingVm请求参数
type GetCandidatePrimaryStoragesForCreatingVmParam struct {
	BaseParam
	Params GetCandidatePrimaryStoragesForCreatingVmDetailParam `json:"params"` // 详细参数
}

