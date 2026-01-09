// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSApplicationPlatformInventoryView SNSApplicationPlatform
type SNSApplicationPlatformInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	State *string `json:"state,omitempty"`
	Type *string `json:"type,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// DeleteSNSApplicationPlatformEventView DeleteSNSApplicationPlatformEvent
type DeleteSNSApplicationPlatformEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangeSNSApplicationPlatformStateEventView ChangeSNSApplicationPlatformStateEvent
type ChangeSNSApplicationPlatformStateEventView struct {
	Inventory SNSApplicationPlatformInventoryView `json:"inventory,omitempty"`
}

// QuerySNSApplicationPlatformView QuerySNSApplicationPlatform
type QuerySNSApplicationPlatformView struct {
	Inventories []SNSApplicationPlatformInventoryView `json:"inventories,omitempty"`
}

