// Copyright (c) ZStack.io, Inc.

package param

// DeleteFirewallIpSetTemplateDetailParam DeleteFirewallIpSetTemplate detail param
type DeleteFirewallIpSetTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteFirewallIpSetTemplateParam DeleteFirewallIpSetTemplate request param
type DeleteFirewallIpSetTemplateParam struct {
	BaseParam
	Params DeleteFirewallIpSetTemplateDetailParam `json:"params"`
}
