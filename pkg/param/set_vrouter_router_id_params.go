// Copyright (c) ZStack.io, Inc.

package param

// SetVRouterRouterIdDetailParam SetVRouterRouterId详细参数
type SetVRouterRouterIdDetailParam struct {
	rest string `json:"vRouterUuid" validate:"required"` // 必填
	rest string `json:"routerId" validate:"required"` // 必填
}

// SetVRouterRouterIdParam SetVRouterRouterId请求参数
type SetVRouterRouterIdParam struct {
	BaseParam
	Params SetVRouterRouterIdDetailParam `json:"params"` // 详细参数
}

