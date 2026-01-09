// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSUniversalSmsEndpointInventoryView SNSUniversalSmsEndpoint
type SNSUniversalSmsEndpointInventoryView struct {
	SmsAccessKeyId *string `json:"smsAccessKeyId,omitempty"`
	SmsAccessKeySecret *string `json:"smsAccessKeySecret,omitempty"`
	Supplier *string `json:"supplier,omitempty"`
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	State *string `json:"state,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	ConnectionStatus *string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

// QuerySNSUniversalSmsEndpointView QuerySNSUniversalSmsEndpoint
type QuerySNSUniversalSmsEndpointView struct {
	Inventories []SNSUniversalSmsEndpointInventoryView `json:"inventories,omitempty"`
}

// UpdateSNSApplicationEndpointEventView UpdateSNSApplicationEndpointEvent
type UpdateSNSApplicationEndpointEventView struct {
	Inventory SNSApplicationEndpointInventoryView `json:"inventory,omitempty"`
}

// CreateSNSUniversalSmsEndpointEventView CreateSNSUniversalSmsEndpointEvent
type CreateSNSUniversalSmsEndpointEventView struct {
	Inventory SNSUniversalSmsEndpointInventoryView `json:"inventory,omitempty"`
}

// ValidateSNSApplicationEndpointEventView ValidateSNSApplicationEndpointEvent
type ValidateSNSApplicationEndpointEventView struct {
	Success bool `json:"success,omitempty"`
}

