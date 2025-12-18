// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EipInventoryView Eip
type EipInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"vmNicUuid,omitempty"`
	rest string `json:"vipUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"vipIp,omitempty"`
	rest string `json:"guestIp,omitempty"`
}

