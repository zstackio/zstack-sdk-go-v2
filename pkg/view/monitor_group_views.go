// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MonitorGroupInventoryView MonitorGroup
type MonitorGroupInventoryView struct {
	Name string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
	Actions string `json:"actions,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	MonitorGroupTemplateRefs []MonitorGroupTemplateRefVOView `json:"monitorGroupTemplateRefs,omitempty"`
}

