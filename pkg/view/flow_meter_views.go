// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// FlowMeterInventoryView FlowMeter
type FlowMeterInventoryView struct {
	Collectors []FlowCollectorInventoryView `json:"collectors,omitempty"`
	NetworkRefs []NetworkRouterFlowMeterRefInventoryView `json:"networkRefs,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Sample int64 `json:"sample,omitempty"`
	ExpireInterval int64 `json:"expireInterval,omitempty"`
	Version string `json:"version,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

