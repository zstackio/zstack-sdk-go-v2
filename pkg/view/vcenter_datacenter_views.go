// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VCenterDatacenterInventoryView VCenterDatacenter
type VCenterDatacenterInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VCenterUuid *string `json:"vCenterUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Morval *string `json:"morval,omitempty"`
}

// QueryVCenterDatacenterView QueryVCenterDatacenter
type QueryVCenterDatacenterView struct {
	Inventories []VCenterDatacenterInventoryView `json:"inventories,omitempty"`
}

