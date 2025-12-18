// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcIkeConfigLocalDetailParam DeleteVpcIkeConfigLocal详细参数
type DeleteVpcIkeConfigLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVpcIkeConfigLocalParam DeleteVpcIkeConfigLocal请求参数
type DeleteVpcIkeConfigLocalParam struct {
	BaseParam
	Params DeleteVpcIkeConfigLocalDetailParam `json:"params"` // 详细参数
}

