// Copyright (c) ZStack.io, Inc.

package param

// RemoveDnsFromVpcRouterDetailParam RemoveDnsFromVpcRouter详细参数
type RemoveDnsFromVpcRouterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"dns" validate:"required"` // 必填
}

// RemoveDnsFromVpcRouterParam RemoveDnsFromVpcRouter请求参数
type RemoveDnsFromVpcRouterParam struct {
	BaseParam
	Params RemoveDnsFromVpcRouterDetailParam `json:"params"` // 详细参数
}

