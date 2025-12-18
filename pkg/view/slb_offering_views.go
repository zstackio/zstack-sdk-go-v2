// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SlbOfferingInventoryView SlbOffering
type SlbOfferingInventoryView struct {
	rest string `json:"managementNetworkUuid,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"imageUuid,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest int `json:"cpuNum,omitempty"`
	rest int `json:"cpuSpeed,omitempty"`
	rest int64 `json:"memorySize,omitempty"`
	rest int64 `json:"reservedMemorySize,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"allocatorStrategy,omitempty"`
	rest int `json:"sortKey,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"state,omitempty"`
}

