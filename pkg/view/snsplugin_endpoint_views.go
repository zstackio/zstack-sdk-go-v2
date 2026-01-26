// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSPluginEndpointInventoryView SNSPluginEndpoint
type SNSPluginEndpointInventoryView struct {
	BaseInfoView
	BaseTimeView
	PluginDriverUuid string `json:"pluginDriverUuid,omitempty"`
	TimeoutInSeconds int64 `json:"timeoutInSeconds,omitempty"`
	Properties interface{} `json:"properties,omitempty"`
	Driver PluginDriverInventoryView `json:"driver,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

// QuerySNSPluginEndpointView QuerySNSPluginEndpoint
type QuerySNSPluginEndpointView struct {
	Inventories []SNSPluginEndpointInventoryView `json:"inventories,omitempty"`
}

// CreateSNSPluginEndpointEventView CreateSNSPluginEndpointEvent
type CreateSNSPluginEndpointEventView struct {
	Inventory SNSPluginEndpointInventoryView `json:"inventory,omitempty"`
}

