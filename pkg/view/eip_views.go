// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EipInventoryView Eip
type EipInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	VmNicUuid string `json:"vmNicUuid,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	State string `json:"state,omitempty"`
	VipIp string `json:"vipIp,omitempty"`
	GuestIp string `json:"guestIp,omitempty"`
}

