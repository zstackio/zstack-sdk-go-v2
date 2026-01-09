// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSMicrosoftTeamsEndpointInventoryView SNSMicrosoftTeamsEndpoint
type SNSMicrosoftTeamsEndpointInventoryView struct {
	Url *string `json:"url,omitempty"`
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

// CreateSNSMicrosoftTeamsEndpointEventView CreateSNSMicrosoftTeamsEndpointEvent
type CreateSNSMicrosoftTeamsEndpointEventView struct {
	Inventory SNSMicrosoftTeamsEndpointInventoryView `json:"inventory,omitempty"`
}

// QuerySNSMicrosoftTeamsEndpointView QuerySNSMicrosoftTeamsEndpoint
type QuerySNSMicrosoftTeamsEndpointView struct {
	Inventories []SNSMicrosoftTeamsEndpointInventoryView `json:"inventories,omitempty"`
}

