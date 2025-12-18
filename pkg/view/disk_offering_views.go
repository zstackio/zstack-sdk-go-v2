// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// DiskOfferingInventoryView DiskOffering
type DiskOfferingInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DiskSize int64 `json:"diskSize,omitempty"`
	SortKey int `json:"sortKey,omitempty"`
	State string `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	AllocatorStrategy string `json:"allocatorStrategy,omitempty"`
}

