// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VniRangeInventoryView VniRange
type VniRangeInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest int `json:"startVni,omitempty"`
	rest int `json:"endVni,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"l2NetworkUuid,omitempty"`
}

