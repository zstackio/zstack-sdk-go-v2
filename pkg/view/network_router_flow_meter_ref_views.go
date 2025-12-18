// Copyright (c) ZStack.io, Inc.

package view

import "time"

// NetworkRouterFlowMeterRefInventoryView NetworkRouterFlowMeterRef
type NetworkRouterFlowMeterRefInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"vRouterUuid,omitempty"`
	rest string `json:"flowMeterUuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

