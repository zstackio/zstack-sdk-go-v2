// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SharedBlockCapacityInventoryView SharedBlockCapacity
type SharedBlockCapacityInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

