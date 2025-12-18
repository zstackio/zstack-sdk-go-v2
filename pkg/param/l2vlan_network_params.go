// Copyright (c) ZStack.io, Inc.

package param

// CreateL2VlanNetworkDetailParam CreateL2VlanNetwork详细参数
type CreateL2VlanNetworkDetailParam struct {
	rest int `json:"vlan" validate:"required"` // 必填
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

// CreateL2VlanNetworkParam CreateL2VlanNetwork请求参数
type CreateL2VlanNetworkParam struct {
	BaseParam
	Params CreateL2VlanNetworkDetailParam `json:"params"` // 详细参数
}

