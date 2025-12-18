// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcFirewallVRouterRefInventoryView VpcFirewallVRouterRef
type VpcFirewallVRouterRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"vpcFirewallUuid,omitempty"`
	rest string `json:"vRouterUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

