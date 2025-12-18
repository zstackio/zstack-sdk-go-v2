// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// FlowCollectorInventoryView FlowCollector
type FlowCollectorInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	FlowMeterUuid string `json:"flowMeterUuid,omitempty"`
	Server string `json:"server,omitempty"`
	Port int64 `json:"port,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

