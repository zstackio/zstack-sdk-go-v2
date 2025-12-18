// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcHaGroupVipRefInventoryView VpcHaGroupVipRef
type VpcHaGroupVipRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"vpcHaRouterUuid,omitempty"`
	rest string `json:"vipUuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"ip,omitempty"`
	rest string `json:"netmask,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

