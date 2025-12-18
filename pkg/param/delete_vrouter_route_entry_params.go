// Copyright (c) ZStack.io, Inc.

package param

// DeleteVRouterRouteEntryDetailParam DeleteVRouterRouteEntry detail param
type DeleteVRouterRouteEntryDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	RouteTableUuid string `json:"routeTableUuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVRouterRouteEntryParam DeleteVRouterRouteEntry request param
type DeleteVRouterRouteEntryParam struct {
	BaseParam
	Params DeleteVRouterRouteEntryDetailParam `json:"params"`
}
