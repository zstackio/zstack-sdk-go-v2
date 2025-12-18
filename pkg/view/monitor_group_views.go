// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MonitorGroupInventoryView MonitorGroup
type MonitorGroupInventoryView struct {
	rest string `json:"name,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"actions,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest []interface{} `json:"monitorGroupTemplateRefs,omitempty"`
}

