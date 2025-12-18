// Copyright (c) ZStack.io, Inc.

package param

// AddHostRouteToL3NetworkDetailParam AddHostRouteToL3Network详细参数
type AddHostRouteToL3NetworkDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"prefix" validate:"required"` // 必填
	rest string `json:"nexthop" validate:"required"` // 必填
}

// AddHostRouteToL3NetworkParam AddHostRouteToL3Network请求参数
type AddHostRouteToL3NetworkParam struct {
	BaseParam
	Params AddHostRouteToL3NetworkDetailParam `json:"params"` // 详细参数
}

