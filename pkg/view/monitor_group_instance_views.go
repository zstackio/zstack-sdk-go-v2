// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MonitorGroupInstanceInventoryView MonitorGroupInstance
type MonitorGroupInstanceInventoryView struct {
	rest string `json:"groupUuid,omitempty"`
	rest string `json:"instanceResourceType,omitempty"`
	rest string `json:"instanceUuid,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"uuid,omitempty"`
}

