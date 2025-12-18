// Copyright (c) ZStack.io, Inc.

package param

// CreateL2NoVlanNetworkDetailParam CreateL2NoVlanNetwork详细参数
type CreateL2NoVlanNetworkDetailParam struct {
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

// CreateL2NoVlanNetworkParam CreateL2NoVlanNetwork请求参数
type CreateL2NoVlanNetworkParam struct {
	BaseParam
	Params CreateL2NoVlanNetworkDetailParam `json:"params"` // 详细参数
}

