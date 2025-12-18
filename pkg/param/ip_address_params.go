// Copyright (c) ZStack.io, Inc.

package param

// DeleteIpAddressDetailParam DeleteIpAddress详细参数
type DeleteIpAddressDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest []string `json:"usedIpUuids" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteIpAddressParam DeleteIpAddress请求参数
type DeleteIpAddressParam struct {
	BaseParam
	Params DeleteIpAddressDetailParam `json:"params"` // 详细参数
}

