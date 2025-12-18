// Copyright (c) ZStack.io, Inc.

package param

// UpdateVRouterRouteTableDetailParam UpdateVRouterRouteTable detail param
type UpdateVRouterRouteTableDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateVRouterRouteTableParam UpdateVRouterRouteTable request param
type UpdateVRouterRouteTableParam struct {
	BaseParam
	Params UpdateVRouterRouteTableDetailParam `json:"params"`
}
