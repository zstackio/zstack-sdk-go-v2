// Copyright (c) ZStack.io, Inc.

package param

// CreateL2VxlanNetworkPoolDetailParam CreateL2VxlanNetworkPool详细参数
type CreateL2VxlanNetworkPoolDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"physicalInterface,omitempty"` // 必填
	rest string `json:"type,omitempty"`
	rest string `json:"vSwitchType,omitempty"`
	rest bool `json:"isolated,omitempty"`
	rest string `json:"pvlan,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateL2VxlanNetworkPoolParam CreateL2VxlanNetworkPool请求参数
type CreateL2VxlanNetworkPoolParam struct {
	BaseParam
	Params CreateL2VxlanNetworkPoolDetailParam `json:"params"` // 详细参数
}

