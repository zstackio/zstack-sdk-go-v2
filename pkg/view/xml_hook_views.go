// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// XmlHookInventoryView XmlHook
type XmlHookInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	HookScript string `json:"hookScript,omitempty"`
	LibvirtVersion string `json:"libvirtVersion,omitempty"`
}

// UpdateVmUserDefinedXmlHookScriptEventView UpdateVmUserDefinedXmlHookScriptEvent
type UpdateVmUserDefinedXmlHookScriptEventView struct {
	Inventory XmlHookInventoryView `json:"inventory,omitempty"`
}

// QueryVmUserDefinedXmlHookScriptView QueryVmUserDefinedXmlHookScript
type QueryVmUserDefinedXmlHookScriptView struct {
	Inventories []XmlHookInventoryView `json:"inventories,omitempty"`
}

// CreateVmUserDefinedXmlHookScriptEventView CreateVmUserDefinedXmlHookScriptEvent
type CreateVmUserDefinedXmlHookScriptEventView struct {
	Inventory XmlHookInventoryView `json:"inventory,omitempty"`
}

