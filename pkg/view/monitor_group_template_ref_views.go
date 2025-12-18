// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MonitorGroupTemplateRefInventoryView MonitorGroupTemplateRef
type MonitorGroupTemplateRefInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"templateUuid,omitempty"`
	rest string `json:"groupUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest bool `json:"isApplied,omitempty"`
}

