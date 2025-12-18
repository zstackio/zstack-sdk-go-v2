// Copyright (c) ZStack.io, Inc.

package param

// UpdateFirewallIpSetTemplateDetailParam UpdateFirewallIpSetTemplate detail param
type UpdateFirewallIpSetTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	SourceValue string `json:"sourceValue,omitempty"`
	DestValue string `json:"destValue,omitempty"`
	Type string `json:"type,omitempty"`
}

// UpdateFirewallIpSetTemplateParam UpdateFirewallIpSetTemplate request param
type UpdateFirewallIpSetTemplateParam struct {
	BaseParam
	Params UpdateFirewallIpSetTemplateDetailParam `json:"params"`
}
