// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcFirewallIpSetTemplateInventoryView VpcFirewallIpSetTemplate
type VpcFirewallIpSetTemplateInventoryView struct {
	Name string `json:"name,omitempty"`
	SourceValue string `json:"sourceValue,omitempty"`
	DestValue string `json:"destValue,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

