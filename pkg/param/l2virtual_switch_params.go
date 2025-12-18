// Copyright (c) ZStack.io, Inc.

package param

// CreateL2VirtualSwitchDetailParam CreateL2VirtualSwitch详细参数
type CreateL2VirtualSwitchDetailParam struct {
	rest bool `json:"isDistributed,omitempty"`
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

// CreateL2VirtualSwitchParam CreateL2VirtualSwitch请求参数
type CreateL2VirtualSwitchParam struct {
	BaseParam
	Params CreateL2VirtualSwitchDetailParam `json:"params"` // 详细参数
}

