// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSSnmpPlatformInventoryView SNSSnmpPlatform
type SNSSnmpPlatformInventoryView struct {
	SnmpAddress string `json:"snmpAddress,omitempty"`
	SnmpPort int `json:"snmpPort,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
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

