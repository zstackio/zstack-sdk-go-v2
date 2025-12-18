// Copyright (c) ZStack.io, Inc.

package param

// CreateObservabilityServerOfferingDetailParam CreateObservabilityServerOffering详细参数
type CreateObservabilityServerOfferingDetailParam struct {
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

// CreateObservabilityServerOfferingParam CreateObservabilityServerOffering请求参数
type CreateObservabilityServerOfferingParam struct {
	BaseParam
	Params CreateObservabilityServerOfferingDetailParam `json:"params"` // 详细参数
}

