// Copyright (c) ZStack.io, Inc.

package param

// SetIpOnHostNetworkInterfaceDetailParam SetIpOnHostNetworkInterface详细参数
type SetIpOnHostNetworkInterfaceDetailParam struct {
	rest string `json:"interfaceUuid" validate:"required"` // 必填
	rest string `json:"ipAddress,omitempty"`
	rest string `json:"netmask,omitempty"`
}

// SetIpOnHostNetworkInterfaceParam SetIpOnHostNetworkInterface请求参数
type SetIpOnHostNetworkInterfaceParam struct {
	BaseParam
	Params SetIpOnHostNetworkInterfaceDetailParam `json:"params"` // 详细参数
}

