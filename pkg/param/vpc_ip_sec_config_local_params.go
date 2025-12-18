// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcIpSecConfigLocalDetailParam DeleteVpcIpSecConfigLocal详细参数
type DeleteVpcIpSecConfigLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVpcIpSecConfigLocalParam DeleteVpcIpSecConfigLocal请求参数
type DeleteVpcIpSecConfigLocalParam struct {
	BaseParam
	Params DeleteVpcIpSecConfigLocalDetailParam `json:"params"` // 详细参数
}

