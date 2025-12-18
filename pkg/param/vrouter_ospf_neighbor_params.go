// Copyright (c) ZStack.io, Inc.

package param

// GetVRouterOspfNeighborDetailParam GetVRouterOspfNeighbor详细参数
type GetVRouterOspfNeighborDetailParam struct {
	rest string `json:"vRouterUuid" validate:"required"` // 必填
}

// GetVRouterOspfNeighborParam GetVRouterOspfNeighbor请求参数
type GetVRouterOspfNeighborParam struct {
	BaseParam
	Params GetVRouterOspfNeighborDetailParam `json:"params"` // 详细参数
}

