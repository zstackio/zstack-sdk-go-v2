// Copyright (c) ZStack.io, Inc.

package param

// CreateOvnControllerOfferingDetailParam CreateOvnControllerOffering详细参数
type CreateOvnControllerOfferingDetailParam struct {
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"managementNetworkUuid" validate:"required"` // 必填
	rest string `json:"imageUuid" validate:"required"` // 必填
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

// CreateOvnControllerOfferingParam CreateOvnControllerOffering请求参数
type CreateOvnControllerOfferingParam struct {
	BaseParam
	Params CreateOvnControllerOfferingDetailParam `json:"params"` // 详细参数
}

