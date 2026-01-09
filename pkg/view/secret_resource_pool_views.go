// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SecretResourcePoolInventoryView SecretResourcePool
type SecretResourcePoolInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	State *string `json:"state,omitempty"`
	Status *string `json:"status,omitempty"`
	Model *string `json:"model,omitempty"`
	HeartbeatInterval *int `json:"heartbeatInterval,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	Ability *string `json:"ability,omitempty"`
}

// DeleteSecretResourcePoolEventView DeleteSecretResourcePoolEvent
type DeleteSecretResourcePoolEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangeSecretResourcePoolStateEventView ChangeSecretResourcePoolStateEvent
type ChangeSecretResourcePoolStateEventView struct {
	Inventory SecretResourcePoolInventoryView `json:"inventory,omitempty"`
}

// QuerySecretResourcePoolView QuerySecretResourcePool
type QuerySecretResourcePoolView struct {
	Inventories []SecretResourcePoolInventoryView `json:"inventories,omitempty"`
}

