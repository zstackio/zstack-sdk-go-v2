// Copyright (c) ZStack.io, Inc.

package view

import "time"

// L2NetworkHostRefInventoryView L2NetworkHostRef
type L2NetworkHostRefInventoryView struct {
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"l2NetworkUuid,omitempty"`
	rest string `json:"l2ProviderType,omitempty"`
	rest string `json:"attachStatus,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

