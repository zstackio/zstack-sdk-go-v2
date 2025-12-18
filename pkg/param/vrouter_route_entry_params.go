// Copyright (c) ZStack.io, Inc.

package param

// DeleteVRouterRouteEntryDetailParam DeleteVRouterRouteEntry详细参数
type DeleteVRouterRouteEntryDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"routeTableUuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVRouterRouteEntryParam DeleteVRouterRouteEntry请求参数
type DeleteVRouterRouteEntryParam struct {
	BaseParam
	Params DeleteVRouterRouteEntryDetailParam `json:"params"` // 详细参数
}

