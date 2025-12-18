// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BareMetal2BondingNicRefInventoryView BareMetal2BondingNicRef
type BareMetal2BondingNicRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"nicUuid,omitempty"`
	rest string `json:"instanceUuid,omitempty"`
	rest string `json:"bondingUuid,omitempty"`
	rest string `json:"provisionNicUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest VmNicInventoryView `json:"vmNic,omitempty"`
	rest BareMetal2InstanceProvisionNicInventoryView `json:"provisionNic,omitempty"`
	rest BareMetal2BondingInventoryView `json:"bareMetal2Bonding,omitempty"`
}

