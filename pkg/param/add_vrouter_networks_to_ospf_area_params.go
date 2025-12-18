// Copyright (c) ZStack.io, Inc.

package param

// AddVRouterNetworksToOspfAreaDetailParam AddVRouterNetworksToOspfArea detail param
type AddVRouterNetworksToOspfAreaDetailParam struct {
	RouterAreaUuid string `json:"routerAreaUuid" validate:"required"`
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddVRouterNetworksToOspfAreaParam AddVRouterNetworksToOspfArea request param
type AddVRouterNetworksToOspfAreaParam struct {
	BaseParam
	Params AddVRouterNetworksToOspfAreaDetailParam `json:"params"`
}
