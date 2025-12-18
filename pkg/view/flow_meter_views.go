// Copyright (c) ZStack.io, Inc.

package view

import "time"

// FlowMeterInventoryView FlowMeter
type FlowMeterInventoryView struct {
	rest []FlowCollectorInventoryView `json:"collectors,omitempty"`
	rest []NetworkRouterFlowMeterRefInventoryView `json:"networkRefs,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest int64 `json:"sample,omitempty"`
	rest int64 `json:"expireInterval,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

