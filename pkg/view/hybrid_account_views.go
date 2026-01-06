// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HybridAccountInventoryView HybridAccount
type HybridAccountInventoryView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	UserUuid string `json:"userUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Akey string `json:"akey,omitempty"`
	HybridAccountId string `json:"hybridAccountId,omitempty"`
	HybridUserId string `json:"hybridUserId,omitempty"`
	HybridUserName string `json:"hybridUserName,omitempty"`
	Current string `json:"current,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryHybridKeySecretView QueryHybridKeySecret
type QueryHybridKeySecretView struct {
	Inventories []HybridAccountInventoryView `json:"inventories,omitempty"`
}

// UpdateAliyunKeySecretEventView UpdateAliyunKeySecretEvent
type UpdateAliyunKeySecretEventView struct {
	Inventory HybridAccountInventoryView `json:"inventory,omitempty"`
}

// AddHybridKeySecretEventView AddHybridKeySecretEvent
type AddHybridKeySecretEventView struct {
	Inventory HybridAccountInventoryView `json:"inventory,omitempty"`
}

// AddAliyunKeySecretEventView AddAliyunKeySecretEvent
type AddAliyunKeySecretEventView struct {
	Inventory HybridAccountInventoryView `json:"inventory,omitempty"`
}

// UpdateHybridKeySecretEventView UpdateHybridKeySecretEvent
type UpdateHybridKeySecretEventView struct {
	Inventory HybridAccountInventoryView `json:"inventory,omitempty"`
}

