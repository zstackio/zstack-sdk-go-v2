// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AccountResourceRefInventoryView AccountResourceRef
type AccountResourceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccountUuid string `json:"accountUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	ConcreteResourceType string `json:"concreteResourceType,omitempty"`
}

// QueryAccountResourceRefView QueryAccountResourceRef
type QueryAccountResourceRefView struct {
	Inventories []AccountResourceRefInventoryView `json:"inventories,omitempty"`
}

// ChangeResourceOwnerEventView ChangeResourceOwnerEvent
type ChangeResourceOwnerEventView struct {
	Inventory AccountResourceRefInventoryView `json:"inventory,omitempty"`
}

