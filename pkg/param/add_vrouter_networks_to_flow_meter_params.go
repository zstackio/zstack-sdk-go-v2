// Copyright (c) ZStack.io, Inc.

package param

// AddVRouterNetworksToFlowMeterDetailParam AddVRouterNetworksToFlowMeter detail param
type AddVRouterNetworksToFlowMeterDetailParam struct {
	FlowMeterUuid string `json:"flowMeterUuid" validate:"required"`
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddVRouterNetworksToFlowMeterParam AddVRouterNetworksToFlowMeter request param
type AddVRouterNetworksToFlowMeterParam struct {
	BaseParam
	Params AddVRouterNetworksToFlowMeterDetailParam `json:"params"`
}
