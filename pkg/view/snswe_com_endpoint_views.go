// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSWeComEndpointInventoryView SNSWeComEndpoint
type SNSWeComEndpointInventoryView struct {
	Url string `json:"url,omitempty"`
	AtAll bool `json:"atAll,omitempty"`
	AtPersonUserIds []string `json:"atPersonUserIds,omitempty"`
	AtPersonList []SNSWeComAtPersonInventoryView `json:"atPersonList,omitempty"`
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

// CreateSNSWeComEndpointEventView CreateSNSWeComEndpointEvent
type CreateSNSWeComEndpointEventView struct {
	Inventory SNSWeComEndpointInventoryView `json:"inventory,omitempty"`
}

// QuerySNSWeComEndpointView QuerySNSWeComEndpoint
type QuerySNSWeComEndpointView struct {
	Inventories []SNSWeComEndpointInventoryView `json:"inventories,omitempty"`
}

