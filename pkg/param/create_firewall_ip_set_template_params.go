// Copyright (c) ZStack.io, Inc.

package param

// CreateFirewallIpSetTemplateDetailParam CreateFirewallIpSetTemplate detail param
type CreateFirewallIpSetTemplateDetailParam struct {
	Name string `json:"name" validate:"required"`
	SourceValue string `json:"sourceValue,omitempty"`
	DestValue string `json:"destValue,omitempty"`
	Type string `json:"type" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFirewallIpSetTemplateParam CreateFirewallIpSetTemplate request param
type CreateFirewallIpSetTemplateParam struct {
	BaseParam
	Params CreateFirewallIpSetTemplateDetailParam `json:"params"`
}
