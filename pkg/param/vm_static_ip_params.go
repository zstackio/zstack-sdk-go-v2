// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmStaticIpDetailParam DeleteVmStaticIp详细参数
type DeleteVmStaticIpDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"staticIp,omitempty"`
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVmStaticIpParam DeleteVmStaticIp请求参数
type DeleteVmStaticIpParam struct {
	BaseParam
	Params DeleteVmStaticIpDetailParam `json:"params"` // 详细参数
}

