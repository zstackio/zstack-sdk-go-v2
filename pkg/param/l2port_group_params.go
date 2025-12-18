// Copyright (c) ZStack.io, Inc.

package param

// CreateL2PortGroupDetailParam CreateL2PortGroup详细参数
type CreateL2PortGroupDetailParam struct {
	rest string `json:"vSwitchUuid" validate:"required"` // 必填
	rest string `json:"vlanMode,omitempty"`
	rest int `json:"vlan" validate:"required"` // 必填
	rest string `json:"vlanRanges,omitempty"`
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

// CreateL2PortGroupParam CreateL2PortGroup请求参数
type CreateL2PortGroupParam struct {
	BaseParam
	Params CreateL2PortGroupDetailParam `json:"params"` // 详细参数
}

