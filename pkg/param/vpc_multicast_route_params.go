// Copyright (c) ZStack.io, Inc.

package param

// GetVpcMulticastRouteDetailParam GetVpcMulticastRoute详细参数
type GetVpcMulticastRouteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVpcMulticastRouteParam GetVpcMulticastRoute请求参数
type GetVpcMulticastRouteParam struct {
	BaseParam
	Params GetVpcMulticastRouteDetailParam `json:"params"` // 详细参数
}

