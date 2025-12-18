// Copyright (c) ZStack.io, Inc.

package param

// CreateConnectionBetweenL3NetworkAndAliyunVSwitchDetailParam CreateConnectionBetweenL3NetworkAndAliyunVSwitch详细参数
type CreateConnectionBetweenL3NetworkAndAliyunVSwitchDetailParam struct {
	rest string `json:"l3networkUuid" validate:"required"` // 必填
	rest string `json:"vpcUuid" validate:"required"` // 必填
	rest string `json:"vbrUuid" validate:"required"` // 必填
	rest string `json:"cpeIp" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"direction" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateConnectionBetweenL3NetworkAndAliyunVSwitchParam CreateConnectionBetweenL3NetworkAndAliyunVSwitch请求参数
type CreateConnectionBetweenL3NetworkAndAliyunVSwitchParam struct {
	BaseParam
	Params CreateConnectionBetweenL3NetworkAndAliyunVSwitchDetailParam `json:"params"` // 详细参数
}

