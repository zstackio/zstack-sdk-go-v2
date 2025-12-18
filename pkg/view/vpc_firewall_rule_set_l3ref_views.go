// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcFirewallRuleSetL3RefInventoryView VpcFirewallRuleSetL3Ref
type VpcFirewallRuleSetL3RefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	RuleSetUuid string `json:"ruleSetUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	VpcFirewallUuid string `json:"vpcFirewallUuid,omitempty"`
	PacketsForwardType string `json:"packetsForwardType,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

