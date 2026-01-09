// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSApplicationEndpointInventoryView SNSApplicationEndpoint
type SNSApplicationEndpointInventoryView struct {
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

// CreateSNSApplicationEndpointEventView CreateSNSApplicationEndpointEvent
type CreateSNSApplicationEndpointEventView struct {
	Inventory SNSApplicationEndpointInventoryView `json:"inventory,omitempty"`
}

// ChangeSNSApplicationEndpointStateEventView ChangeSNSApplicationEndpointStateEvent
type ChangeSNSApplicationEndpointStateEventView struct {
	Inventory SNSApplicationEndpointInventoryView `json:"inventory,omitempty"`
}

// QuerySNSApplicationEndpointView QuerySNSApplicationEndpoint
type QuerySNSApplicationEndpointView struct {
	Inventories []SNSApplicationEndpointInventoryView `json:"inventories,omitempty"`
}

// DeleteSNSApplicationEndpointEventView DeleteSNSApplicationEndpointEvent
type DeleteSNSApplicationEndpointEventView struct {
	Success bool `json:"success,omitempty"`
}

