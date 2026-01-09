// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// RouterAreaInventoryView RouterArea
type RouterAreaInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AreaId *string `json:"areaId,omitempty"`
	Type *string `json:"type,omitempty"`
	Authentication *string `json:"authentication,omitempty"`
	Password *string `json:"password,omitempty"`
	KeyId *int `json:"keyId,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// UpdateVRouterOspfAreaEventView UpdateVRouterOspfAreaEvent
type UpdateVRouterOspfAreaEventView struct {
	Inventory RouterAreaInventoryView `json:"inventory,omitempty"`
}

// CreateVRouterOspfAreaEventView CreateVRouterOspfAreaEvent
type CreateVRouterOspfAreaEventView struct {
	Inventory RouterAreaInventoryView `json:"inventory,omitempty"`
}

// QueryVRouterOspfAreaView QueryVRouterOspfArea
type QueryVRouterOspfAreaView struct {
	Inventories []RouterAreaInventoryView `json:"inventories,omitempty"`
}

