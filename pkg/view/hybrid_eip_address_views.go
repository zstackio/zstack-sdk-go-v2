// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HybridEipAddressInventoryView HybridEipAddress
type HybridEipAddressInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"eipId,omitempty"`
	rest string `json:"bandWidth,omitempty"`
	rest string `json:"dataCenterUuid,omitempty"`
	rest string `json:"allocateResourceUuid,omitempty"`
	rest string `json:"allocateResourceType,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"eipAddress,omitempty"`
	rest string `json:"eipType,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"chargeType,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"allocateTime,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

