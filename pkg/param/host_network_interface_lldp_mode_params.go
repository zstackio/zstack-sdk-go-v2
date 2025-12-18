// Copyright (c) ZStack.io, Inc.

package param

// ChangeHostNetworkInterfaceLldpModeDetailParam ChangeHostNetworkInterfaceLldpMode详细参数
type ChangeHostNetworkInterfaceLldpModeDetailParam struct {
	rest []string `json:"interfaceUuids" validate:"required"` // 必填
	rest string `json:"mode,omitempty"`
}

// ChangeHostNetworkInterfaceLldpModeParam ChangeHostNetworkInterfaceLldpMode请求参数
type ChangeHostNetworkInterfaceLldpModeParam struct {
	BaseParam
	Params ChangeHostNetworkInterfaceLldpModeDetailParam `json:"params"` // 详细参数
}

