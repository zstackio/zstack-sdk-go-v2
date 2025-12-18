// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ResourceStackVmPortRefInventoryView ResourceStackVmPortRef
type ResourceStackVmPortRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"stackUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

