// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcFirewallIpSetTemplateInventoryView VpcFirewallIpSetTemplate
type VpcFirewallIpSetTemplateInventoryView struct {
	rest string `json:"name,omitempty"`
	rest string `json:"sourceValue,omitempty"`
	rest string `json:"destValue,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"uuid,omitempty"`
}

