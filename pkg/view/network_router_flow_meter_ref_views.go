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
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

