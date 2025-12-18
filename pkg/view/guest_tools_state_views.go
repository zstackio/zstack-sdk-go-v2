// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// GuestToolsStateInventoryView GuestToolsState
type GuestToolsStateInventoryView struct {
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	QgaState string `json:"qgaState,omitempty"`
	ZwatchState string `json:"zwatchState,omitempty"`
	Version string `json:"version,omitempty"`
	Platform string `json:"platform,omitempty"`
	OsType string `json:"osType,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

