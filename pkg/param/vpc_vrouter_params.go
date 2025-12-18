// Copyright (c) ZStack.io, Inc.

package param

// CreateVpcVRouterDetailParam CreateVpcVRouter详细参数
type CreateVpcVRouterDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"virtualRouterOfferingUuid" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"primaryStorageUuidForRootVolume,omitempty"`
	rest []string `json:"rootVolumeSystemTags,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVpcVRouterParam CreateVpcVRouter请求参数
type CreateVpcVRouterParam struct {
	BaseParam
	Params CreateVpcVRouterDetailParam `json:"params"` // 详细参数
}

