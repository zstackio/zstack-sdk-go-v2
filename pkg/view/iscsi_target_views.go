// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IscsiTargetInventoryView IscsiTarget
type IscsiTargetInventoryView struct {
	rest string `json:"iscsiServerUuid,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"iqn,omitempty"`
	rest []IscsiLunInventoryView `json:"iscsiLuns,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

