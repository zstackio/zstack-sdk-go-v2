// Copyright (c) ZStack.io, Inc.

package param

// CreateL2HardwareVxlanNetworkPoolDetailParam CreateL2HardwareVxlanNetworkPool详细参数
type CreateL2HardwareVxlanNetworkPoolDetailParam struct {
	rest string `json:"sdnControllerUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"physicalInterface" validate:"required"` // 必填
	rest string `json:"type,omitempty"`
	rest string `json:"vSwitchType,omitempty"`
	rest bool `json:"isolated,omitempty"`
	rest string `json:"pvlan,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateL2HardwareVxlanNetworkPoolParam CreateL2HardwareVxlanNetworkPool请求参数
type CreateL2HardwareVxlanNetworkPoolParam struct {
	BaseParam
	Params CreateL2HardwareVxlanNetworkPoolDetailParam `json:"params"` // 详细参数
}

