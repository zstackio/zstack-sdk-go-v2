// Copyright (c) ZStack.io, Inc.

package param

// SetVRouterRouterIdDetailParam SetVRouterRouterId detail param
type SetVRouterRouterIdDetailParam struct {
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
	RouterId string `json:"routerId" validate:"required"`
}

// SetVRouterRouterIdParam SetVRouterRouterId request param
type SetVRouterRouterIdParam struct {
	BaseParam
	Params SetVRouterRouterIdDetailParam `json:"params"`
}
