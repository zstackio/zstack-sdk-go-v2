// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CSPSecretResourcePoolInventoryView CSPSecretResourcePool
type CSPSecretResourcePoolInventoryView struct {
	BaseInfoView
	BaseTimeView
	ManagementIp string `json:"managementIp,omitempty"`
	Port int `json:"port,omitempty"`
	AppId string `json:"appId,omitempty"`
	AppKey string `json:"appKey,omitempty"`
	KeyId string `json:"keyId,omitempty"`
	UserId string `json:"userId,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
	Ability string `json:"ability,omitempty"`
}

// UpdateSecretResourcePoolEventView UpdateSecretResourcePoolEvent
type UpdateSecretResourcePoolEventView struct {
	Inventory SecretResourcePoolInventoryView `json:"inventory,omitempty"`
}

// CreateSecretResourcePoolEventView CreateSecretResourcePoolEvent
type CreateSecretResourcePoolEventView struct {
	Inventory SecretResourcePoolInventoryView `json:"inventory,omitempty"`
}

