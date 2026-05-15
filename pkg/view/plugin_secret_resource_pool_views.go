// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PluginSecretResourcePoolInventoryView PluginSecretResourcePool
type PluginSecretResourcePoolInventoryView struct {
	BaseInfoView
	BaseTimeView
	Properties string `json:"properties,omitempty"`
	PluginDriverUuid string `json:"pluginDriverUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
	Ability string `json:"ability,omitempty"`
}

// CreateSecretResourcePoolEventView CreateSecretResourcePoolEvent
type CreateSecretResourcePoolEventView struct {
	Inventory SecretResourcePoolInventoryView `json:"inventory,omitempty"`
}

// UpdateSecretResourcePoolEventView UpdateSecretResourcePoolEvent
type UpdateSecretResourcePoolEventView struct {
	Inventory SecretResourcePoolInventoryView `json:"inventory,omitempty"`
}

