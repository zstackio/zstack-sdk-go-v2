// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSPluginEndpointInventoryView SNSPluginEndpoint
type SNSPluginEndpointInventoryView struct {
	PluginDriverUuid string `json:"pluginDriverUuid,omitempty"`
	TimeoutInSeconds int64 `json:"timeoutInSeconds,omitempty"`
	Properties interface{} `json:"properties,omitempty"`
	Driver PluginDriverInventoryView `json:"driver,omitempty"`
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

