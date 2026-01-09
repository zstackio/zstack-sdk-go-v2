// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunNasAccessGroupInventoryView AliyunNasAccessGroup
type AliyunNasAccessGroupInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	DataCenterUuid *string `json:"dataCenterUuid,omitempty"`
	Rules []AliyunNasAccessRuleInventoryView `json:"rules,omitempty"`
	Type *string `json:"type,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// AddAliyunNasAccessGroupEventView AddAliyunNasAccessGroupEvent
type AddAliyunNasAccessGroupEventView struct {
	Inventory AliyunNasAccessGroupInventoryView `json:"inventory,omitempty"`
}

// UpdateAliyunNasAccessGroupEventView UpdateAliyunNasAccessGroupEvent
type UpdateAliyunNasAccessGroupEventView struct {
	Inventory AliyunNasAccessGroupInventoryView `json:"inventory,omitempty"`
}

// CreateAliyunNasAccessGroupEventView CreateAliyunNasAccessGroupEvent
type CreateAliyunNasAccessGroupEventView struct {
	Inventory AliyunNasAccessGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteAliyunNasAccessGroupEventView DeleteAliyunNasAccessGroupEvent
type DeleteAliyunNasAccessGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryAliyunNasAccessGroupView QueryAliyunNasAccessGroup
type QueryAliyunNasAccessGroupView struct {
	Inventories []AliyunNasAccessGroupInventoryView `json:"inventories,omitempty"`
}

