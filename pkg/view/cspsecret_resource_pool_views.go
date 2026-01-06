// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// CSPSecretResourcePoolInventoryView CSPSecretResourcePool
type CSPSecretResourcePoolInventoryView struct {
	ManagementIp string `json:"managementIp,omitempty"`
	Port int `json:"port,omitempty"`
	AppId string `json:"appId,omitempty"`
	AppKey string `json:"appKey,omitempty"`
	KeyId string `json:"keyId,omitempty"`
	UserId string `json:"userId,omitempty"`
	Protocol string `json:"protocol,omitempty"`
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

