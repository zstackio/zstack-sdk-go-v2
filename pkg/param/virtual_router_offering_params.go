// Copyright (c) ZStack.io, Inc.

package param

// CreateVirtualRouterOfferingDetailParam CreateVirtualRouterOffering详细参数
type CreateVirtualRouterOfferingDetailParam struct {
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"managementNetworkUuid" validate:"required"` // 必填
	rest string `json:"imageUuid" validate:"required"` // 必填
	rest string `json:"publicNetworkUuid,omitempty"`
	rest bool `json:"isDefault,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest int `json:"cpuNum" validate:"required"` // 必填
	rest int64 `json:"memorySize" validate:"required"` // 必填
	rest int64 `json:"reservedMemorySize,omitempty"`
	rest string `json:"allocatorStrategy,omitempty"`
	rest int `json:"sortKey,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVirtualRouterOfferingParam CreateVirtualRouterOffering请求参数
type CreateVirtualRouterOfferingParam struct {
	BaseParam
	Params CreateVirtualRouterOfferingDetailParam `json:"params"` // 详细参数
}

