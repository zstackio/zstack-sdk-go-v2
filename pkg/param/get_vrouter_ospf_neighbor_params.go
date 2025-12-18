// Copyright (c) ZStack.io, Inc.

package param

// GetVRouterOspfNeighborDetailParam GetVRouterOspfNeighbor detail param
type GetVRouterOspfNeighborDetailParam struct {
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
}

// GetVRouterOspfNeighborParam GetVRouterOspfNeighbor request param
type GetVRouterOspfNeighborParam struct {
	BaseParam
	Params GetVRouterOspfNeighborDetailParam `json:"params"`
}
