// Copyright (c) ZStack.io, Inc.

package param

// GetVRouterRouterIdDetailParam GetVRouterRouterId详细参数
type GetVRouterRouterIdDetailParam struct {
	rest string `json:"vRouterUuid" validate:"required"` // 必填
}

// GetVRouterRouterIdParam GetVRouterRouterId请求参数
type GetVRouterRouterIdParam struct {
	BaseParam
	Params GetVRouterRouterIdDetailParam `json:"params"` // 详细参数
}

