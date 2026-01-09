// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HuaweiIMasterFabricInventoryView HuaweiIMasterFabric
type HuaweiIMasterFabricInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	SdnControllerUuid *string `json:"sdnControllerUuid,omitempty"`
	State *string `json:"state,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// DeleteHuaweiIMasterFabricEventView DeleteHuaweiIMasterFabricEvent
type DeleteHuaweiIMasterFabricEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryHuaweiIMasterFabricView QueryHuaweiIMasterFabric
type QueryHuaweiIMasterFabricView struct {
	Inventories []HuaweiIMasterFabricInventoryView `json:"inventories,omitempty"`
}

