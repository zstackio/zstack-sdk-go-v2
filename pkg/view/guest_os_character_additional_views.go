// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// GuestOsCharacterInventoryView GuestOsCharacter
type GuestOsCharacterInventoryView struct {
	BaseInfoView
	BaseTimeView
	Architecture *string `json:"architecture,omitempty"`
	Platform *string `json:"platform,omitempty"`
	OsRelease *string `json:"osRelease,omitempty"`
	Acpi *bool `json:"acpi,omitempty"`
	HygonTag *bool `json:"hygonTag,omitempty"`
	X2apic *bool `json:"x2apic,omitempty"`
	CpuModel *string `json:"cpuModel,omitempty"`
	NicDriver *string `json:"nicDriver,omitempty"`
}

