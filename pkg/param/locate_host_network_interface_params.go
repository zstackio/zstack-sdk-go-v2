// Copyright (c) ZStack.io, Inc.

package param

// LocateHostNetworkInterfaceDetailParam LocateHostNetworkInterface详细参数
type LocateHostNetworkInterfaceDetailParam struct {
	rest string `json:"hostUuid" validate:"required"` // 必填
	rest string `json:"networkInterfaceName" validate:"required"` // 必填
	rest int64 `json:"interval,omitempty"`
}

// LocateHostNetworkInterfaceParam LocateHostNetworkInterface请求参数
type LocateHostNetworkInterfaceParam struct {
	BaseParam
	Params LocateHostNetworkInterfaceDetailParam `json:"params"` // 详细参数
}

