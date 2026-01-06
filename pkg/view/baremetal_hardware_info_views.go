// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BaremetalHardwareInfoInventoryView BaremetalHardwareInfo
type BaremetalHardwareInfoInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ChassisUuid string `json:"chassisUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

