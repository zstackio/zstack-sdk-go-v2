// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSUniversalSmsEndpointInventoryView SNSUniversalSmsEndpoint
type SNSUniversalSmsEndpointInventoryView struct {
	BaseInfoView
	BaseTimeView
	SmsAccessKeyId string `json:"smsAccessKeyId,omitempty"`
	SmsAccessKeySecret string `json:"smsAccessKeySecret,omitempty"`
	Supplier string `json:"supplier,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

// QuerySNSUniversalSmsEndpointView QuerySNSUniversalSmsEndpoint
type QuerySNSUniversalSmsEndpointView struct {
	Inventories []SNSUniversalSmsEndpointInventoryView `json:"inventories,omitempty"`
}

// CreateSNSUniversalSmsEndpointEventView CreateSNSUniversalSmsEndpointEvent
type CreateSNSUniversalSmsEndpointEventView struct {
	Inventory SNSUniversalSmsEndpointInventoryView `json:"inventory,omitempty"`
}

// ValidateSNSApplicationEndpointEventView ValidateSNSApplicationEndpointEvent
type ValidateSNSApplicationEndpointEventView struct {
	Success bool `json:"success,omitempty"`
}

