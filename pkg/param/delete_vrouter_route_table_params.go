// Copyright (c) ZStack.io, Inc.

package param

// DeleteVRouterRouteTableDetailParam DeleteVRouterRouteTable detail param
type DeleteVRouterRouteTableDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVRouterRouteTableParam DeleteVRouterRouteTable request param
type DeleteVRouterRouteTableParam struct {
	BaseParam
	Params DeleteVRouterRouteTableDetailParam `json:"params"`
}
