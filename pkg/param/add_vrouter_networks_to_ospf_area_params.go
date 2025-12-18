// Copyright (c) ZStack.io, Inc.

package param

// AddVRouterNetworksToOspfAreaDetailParam AddVRouterNetworksToOspfArea详细参数
type AddVRouterNetworksToOspfAreaDetailParam struct {
	rest string `json:"routerAreaUuid" validate:"required"` // 必填
	rest string `json:"vRouterUuid" validate:"required"` // 必填
	rest []string `json:"l3NetworkUuids" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddVRouterNetworksToOspfAreaParam AddVRouterNetworksToOspfArea请求参数
type AddVRouterNetworksToOspfAreaParam struct {
	BaseParam
	Params AddVRouterNetworksToOspfAreaDetailParam `json:"params"` // 详细参数
}

