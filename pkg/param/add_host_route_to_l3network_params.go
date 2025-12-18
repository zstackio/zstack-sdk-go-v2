// Copyright (c) ZStack.io, Inc.

package param

// AddHostRouteToL3NetworkDetailParam AddHostRouteToL3Network detail param
type AddHostRouteToL3NetworkDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Prefix string `json:"prefix" validate:"required"`
	Nexthop string `json:"nexthop" validate:"required"`
}

// AddHostRouteToL3NetworkParam AddHostRouteToL3Network request param
type AddHostRouteToL3NetworkParam struct {
	BaseParam
	Params AddHostRouteToL3NetworkDetailParam `json:"params"`
}
