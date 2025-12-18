// Copyright (c) ZStack.io, Inc.

package param

// GetVRouterRouterIdDetailParam GetVRouterRouterId detail param
type GetVRouterRouterIdDetailParam struct {
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
}

// GetVRouterRouterIdParam GetVRouterRouterId request param
type GetVRouterRouterIdParam struct {
	BaseParam
	Params GetVRouterRouterIdDetailParam `json:"params"`
}
