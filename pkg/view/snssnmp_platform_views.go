// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSSnmpPlatformInventoryView SNSSnmpPlatform
type SNSSnmpPlatformInventoryView struct {
	BaseInfoView
	BaseTimeView
	SnmpAddress string `json:"snmpAddress,omitempty"`
	SnmpPort    int    `json:"snmpPort,omitempty"`
	State       string `json:"state,omitempty"`
	Type        string `json:"type,omitempty"`
}

// QuerySNSEmailPlatformView QuerySNSEmailPlatform
type QuerySNSEmailPlatformView struct {
	Inventories []SNSEmailPlatformInventoryView `json:"inventories,omitempty"`
}

// UpdateSNSApplicationPlatformEventView UpdateSNSApplicationPlatformEvent
type UpdateSNSApplicationPlatformEventView struct {
	Inventory SNSApplicationPlatformInventoryView `json:"inventory,omitempty"`
}

// CreateSNSApplicationPlatformEventView CreateSNSApplicationPlatformEvent
type CreateSNSApplicationPlatformEventView struct {
	Inventory SNSApplicationPlatformInventoryView `json:"inventory,omitempty"`
}
