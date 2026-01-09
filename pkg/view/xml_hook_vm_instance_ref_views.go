// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// XmlHookVmInstanceRefInventoryView XmlHookVmInstanceRef
type XmlHookVmInstanceRefInventoryView struct {
	Id *int64 `json:"id,omitempty"`
	XmlHookUuid *string `json:"xmlHookUuid,omitempty"`
	VmInstanceUuid *string `json:"vmInstanceUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

