// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PortForwardingRuleInventoryView PortForwardingRule
type PortForwardingRuleInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"vipIp,omitempty"`
	rest string `json:"guestIp,omitempty"`
	rest string `json:"vipUuid,omitempty"`
	rest int `json:"vipPortStart,omitempty"`
	rest int `json:"vipPortEnd,omitempty"`
	rest int `json:"privatePortStart,omitempty"`
	rest int `json:"privatePortEnd,omitempty"`
	rest string `json:"vmNicUuid,omitempty"`
	rest string `json:"protocolType,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"allowedCidr,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

