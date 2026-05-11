// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SanSecSecretResourcePoolInventoryView SanSecSecretResourcePool
type SanSecSecretResourcePoolInventoryView struct {
	BaseInfoView
	BaseTimeView
	KeyIndex int `json:"keyIndex,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Port int `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Sm3Key string `json:"sm3Key,omitempty"`
	Sm4Key string `json:"sm4Key,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
	Ability string `json:"ability,omitempty"`
}

