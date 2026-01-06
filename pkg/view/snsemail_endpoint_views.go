// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSEmailEndpointInventoryView SNSEmailEndpoint
type SNSEmailEndpointInventoryView struct {
	Email string `json:"email,omitempty"`
	EmailAddresses []SNSEmailAddressInventoryView `json:"emailAddresses,omitempty"`
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

// QuerySNSEmailEndpointView QuerySNSEmailEndpoint
type QuerySNSEmailEndpointView struct {
	Inventories []SNSEmailEndpointInventoryView `json:"inventories,omitempty"`
}

