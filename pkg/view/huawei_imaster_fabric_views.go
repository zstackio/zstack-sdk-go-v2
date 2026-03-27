// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HuaweiIMasterFabricInventoryView HuaweiIMasterFabric
type HuaweiIMasterFabricInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	State string `json:"state,omitempty"`
}

// DeleteHuaweiIMasterFabricEventView DeleteHuaweiIMasterFabricEvent
type DeleteHuaweiIMasterFabricEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryHuaweiIMasterFabricView QueryHuaweiIMasterFabric
type QueryHuaweiIMasterFabricView struct {
	Inventories []HuaweiIMasterFabricInventoryView `json:"inventories,omitempty"`
}

