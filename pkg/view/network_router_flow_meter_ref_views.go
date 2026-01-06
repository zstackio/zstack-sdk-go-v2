// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NetworkRouterFlowMeterRefInventoryView NetworkRouterFlowMeterRef
type NetworkRouterFlowMeterRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VRouterUuid string `json:"vRouterUuid,omitempty"`
	FlowMeterUuid string `json:"flowMeterUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryVRouterFlowMeterNetworkView QueryVRouterFlowMeterNetwork
type QueryVRouterFlowMeterNetworkView struct {
	Inventories []NetworkRouterFlowMeterRefInventoryView `json:"inventories,omitempty"`
}

// AddVRouterNetworksToFlowMeterEventView AddVRouterNetworksToFlowMeterEvent
type AddVRouterNetworksToFlowMeterEventView struct {
	Inventories []NetworkRouterFlowMeterRefInventoryView `json:"inventories,omitempty"`
}

