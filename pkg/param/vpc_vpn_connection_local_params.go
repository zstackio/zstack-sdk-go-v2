// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcVpnConnectionLocalDetailParam DeleteVpcVpnConnectionLocal详细参数
type DeleteVpcVpnConnectionLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVpcVpnConnectionLocalParam DeleteVpcVpnConnectionLocal请求参数
type DeleteVpcVpnConnectionLocalParam struct {
	BaseParam
	Params DeleteVpcVpnConnectionLocalDetailParam `json:"params"` // 详细参数
}

