// Copyright (c) ZStack.io, Inc.

package param

// AddVRouterRouteEntryDetailParam AddVRouterRouteEntry详细参数
type AddVRouterRouteEntryDetailParam struct {
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"routeTableUuid" validate:"required"` // 必填
	rest string `json:"destination" validate:"required"` // 必填
	rest string `json:"target,omitempty"`
	rest int `json:"distance,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddVRouterRouteEntryParam AddVRouterRouteEntry请求参数
type AddVRouterRouteEntryParam struct {
	BaseParam
	Params AddVRouterRouteEntryDetailParam `json:"params"` // 详细参数
}

