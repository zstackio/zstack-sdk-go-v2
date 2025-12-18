// Copyright (c) ZStack.io, Inc.

package param

// RemoveHostRouteFromL3NetworkDetailParam RemoveHostRouteFromL3Network detail param
type RemoveHostRouteFromL3NetworkDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Prefix string `json:"prefix" validate:"required"`
}

// RemoveHostRouteFromL3NetworkParam RemoveHostRouteFromL3Network request param
type RemoveHostRouteFromL3NetworkParam struct {
	BaseParam
	Params RemoveHostRouteFromL3NetworkDetailParam `json:"params"`
}
