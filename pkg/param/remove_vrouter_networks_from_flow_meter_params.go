// Copyright (c) ZStack.io, Inc.

package param

// RemoveVRouterNetworksFromFlowMeterDetailParam RemoveVRouterNetworksFromFlowMeter detail param
type RemoveVRouterNetworksFromFlowMeterDetailParam struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RemoveVRouterNetworksFromFlowMeterParam RemoveVRouterNetworksFromFlowMeter request param
type RemoveVRouterNetworksFromFlowMeterParam struct {
	BaseParam
	Params RemoveVRouterNetworksFromFlowMeterDetailParam `json:"params"`
}
