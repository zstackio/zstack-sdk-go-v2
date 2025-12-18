// Copyright (c) ZStack.io, Inc.

package param

// CreateL2HardwareVxlanNetworkDetailParam CreateL2HardwareVxlanNetwork详细参数
type CreateL2HardwareVxlanNetworkDetailParam struct {
	rest int `json:"vni,omitempty"`
	rest string `json:"poolUuid" validate:"required"` // 必填
	rest string `json:"h3cTenantUuid,omitempty"`
	rest int `json:"vlan,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid,omitempty"` // 必填
	rest string `json:"physicalInterface,omitempty"` // 必填
	rest string `json:"type,omitempty"`
	rest string `json:"vSwitchType,omitempty"`
	rest bool `json:"isolated,omitempty"`
	rest string `json:"pvlan,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateL2HardwareVxlanNetworkParam CreateL2HardwareVxlanNetwork请求参数
type CreateL2HardwareVxlanNetworkParam struct {
	BaseParam
	Params CreateL2HardwareVxlanNetworkDetailParam `json:"params"` // 详细参数
}

