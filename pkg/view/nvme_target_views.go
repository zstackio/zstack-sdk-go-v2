// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NvmeTargetInventoryView NvmeTarget
type NvmeTargetInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Nqn *string `json:"nqn,omitempty"`
	NvmeServerUuid *string `json:"nvmeServerUuid,omitempty"`
	State *string `json:"state,omitempty"`
	NvmeLuns []NvmeLunInventoryView `json:"nvmeLuns,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// QueryNvmeTargetView QueryNvmeTarget
type QueryNvmeTargetView struct {
	Inventories []NvmeTargetInventoryView `json:"inventories,omitempty"`
}

// RefreshNvmeTargetEventView RefreshNvmeTargetEvent
type RefreshNvmeTargetEventView struct {
	Inventories []NvmeTargetInventoryView `json:"inventories,omitempty"`
}

