// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AiSiNoSecretResourcePoolInventoryView AiSiNoSecretResourcePool
type AiSiNoSecretResourcePoolInventoryView struct {
	BaseInfoView
	BaseTimeView
	ManagementIp      string `json:"managementIp,omitempty"`
	Port              int    `json:"port,omitempty"`
	Route             string `json:"route,omitempty"`
	ClientID          string `json:"clientID,omitempty"`
	ClientSecrete     string `json:"clientSecrete,omitempty"`
	AppId             string `json:"appId,omitempty"`
	KeyNumSM2         string `json:"keyNumSM2,omitempty"`
	KeyNumSM4         string `json:"keyNumSM4,omitempty"`
	ZoneUuid          string `json:"zoneUuid,omitempty"`
	Type              string `json:"type,omitempty"`
	State             string `json:"state,omitempty"`
	Status            string `json:"status,omitempty"`
	Model             string `json:"model,omitempty"`
	HeartbeatInterval int    `json:"heartbeatInterval,omitempty"`
	Ability           string `json:"ability,omitempty"`
}

// CreateSecretResourcePoolEventView CreateSecretResourcePoolEvent
type CreateSecretResourcePoolEventView struct {
	Inventory SecretResourcePoolInventoryView `json:"inventory,omitempty"`
}
