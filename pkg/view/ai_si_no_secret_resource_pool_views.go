// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AiSiNoSecretResourcePoolInventoryView AiSiNoSecretResourcePool
type AiSiNoSecretResourcePoolInventoryView struct {
	ManagementIp string `json:"managementIp,omitempty"`
	Port int `json:"port,omitempty"`
	Route string `json:"route,omitempty"`
	ClientID string `json:"clientID,omitempty"`
	ClientSecrete string `json:"clientSecrete,omitempty"`
	AppId string `json:"appId,omitempty"`
	KeyNumSM2 string `json:"keyNumSM2,omitempty"`
	KeyNumSM4 string `json:"keyNumSM4,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	Ability string `json:"ability,omitempty"`
}

// CreateSecretResourcePoolEventView CreateSecretResourcePoolEvent
type CreateSecretResourcePoolEventView struct {
	Inventory SecretResourcePoolInventoryView `json:"inventory,omitempty"`
}

