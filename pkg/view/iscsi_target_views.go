// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IscsiTargetInventoryView IscsiTarget
type IscsiTargetInventoryView struct {
	IscsiServerUuid string `json:"iscsiServerUuid,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Iqn string `json:"iqn,omitempty"`
	IscsiLuns []IscsiLunInventoryView `json:"iscsiLuns,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

