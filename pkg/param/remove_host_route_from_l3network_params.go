// Copyright (c) ZStack.io, Inc.

package param

// RemoveHostRouteFromL3NetworkDetailParam RemoveHostRouteFromL3Network详细参数
type RemoveHostRouteFromL3NetworkDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"prefix" validate:"required"` // 必填
}

// RemoveHostRouteFromL3NetworkParam RemoveHostRouteFromL3Network请求参数
type RemoveHostRouteFromL3NetworkParam struct {
	BaseParam
	Params RemoveHostRouteFromL3NetworkDetailParam `json:"params"` // 详细参数
}

