// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsVpcInLocalDetailParam DeleteEcsVpcInLocal详细参数
type DeleteEcsVpcInLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteEcsVpcInLocalParam DeleteEcsVpcInLocal请求参数
type DeleteEcsVpcInLocalParam struct {
	BaseParam
	Params DeleteEcsVpcInLocalDetailParam `json:"params"` // 详细参数
}

