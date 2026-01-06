// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SanSecSecretResourcePoolInventoryView SanSecSecretResourcePool
type SanSecSecretResourcePoolInventoryView struct {
	KeyIndex int `json:"keyIndex,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Port int `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Sm3Key string `json:"sm3Key,omitempty"`
	Sm4Key string `json:"sm4Key,omitempty"`
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

