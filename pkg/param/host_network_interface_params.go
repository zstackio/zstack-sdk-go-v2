// Copyright (c) ZStack.io, Inc.

package param

// UpdateHostNetworkInterfaceDetailParam UpdateHostNetworkInterface详细参数
type UpdateHostNetworkInterfaceDetailParam struct {
	rest string `json:"interfaceUuid" validate:"required"` // 必填
	rest string `json:"description" validate:"required"` // 必填
}

// UpdateHostNetworkInterfaceParam UpdateHostNetworkInterface请求参数
type UpdateHostNetworkInterfaceParam struct {
	BaseParam
	Params UpdateHostNetworkInterfaceDetailParam `json:"params"` // 详细参数
}

