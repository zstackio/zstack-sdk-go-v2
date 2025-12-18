// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcFirewallRuleSetL3RefInventoryView VpcFirewallRuleSetL3Ref
type VpcFirewallRuleSetL3RefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"ruleSetUuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"vpcFirewallUuid,omitempty"`
	rest string `json:"packetsForwardType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

