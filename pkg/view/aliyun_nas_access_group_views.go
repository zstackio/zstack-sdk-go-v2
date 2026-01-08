// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunNasAccessGroupInventoryView AliyunNasAccessGroup
type AliyunNasAccessGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	DataCenterUuid string                             `json:"dataCenterUuid,omitempty"`
	Rules          []AliyunNasAccessRuleInventoryView `json:"rules,omitempty"`
	Type           string                             `json:"type,omitempty"`
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
