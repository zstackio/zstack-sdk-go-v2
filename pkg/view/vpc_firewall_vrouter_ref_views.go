// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcFirewallVRouterRefInventoryView VpcFirewallVRouterRef
type VpcFirewallVRouterRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	VpcFirewallUuid string `json:"vpcFirewallUuid,omitempty"`
	VRouterUuid string `json:"vRouterUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

