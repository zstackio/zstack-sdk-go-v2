// Copyright (c) ZStack.io, Inc.

package view

import "time"

// XmlHookVmInstanceRefInventoryView XmlHookVmInstanceRef
type XmlHookVmInstanceRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"xmlHookUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

