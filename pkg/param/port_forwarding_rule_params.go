// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeletePortForwardingRuleParamDetail DeletePortForwardingRule detail param
type DeletePortForwardingRuleParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePortForwardingRuleParam DeletePortForwardingRule request param
type DeletePortForwardingRuleParam struct {
	BaseParam
	DeletePortForwardingRule DeletePortForwardingRuleParamDetail `json:"deletePortForwardingRule"`
}
// UpdatePortForwardingRuleParamDetail UpdatePortForwardingRule detail param
type UpdatePortForwardingRuleParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdatePortForwardingRuleParam UpdatePortForwardingRule request param
type UpdatePortForwardingRuleParam struct {
	BaseParam
	UpdatePortForwardingRule UpdatePortForwardingRuleParamDetail `json:"updatePortForwardingRule"`
}
// DetachPortForwardingRuleParamDetail DetachPortForwardingRule detail param
type DetachPortForwardingRuleParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DetachPortForwardingRuleParam DetachPortForwardingRule request param
type DetachPortForwardingRuleParam struct {
	BaseParam
	DetachPortForwardingRule DetachPortForwardingRuleParamDetail `json:"detachPortForwardingRule"`
}
// AttachPortForwardingRuleParamDetail AttachPortForwardingRule detail param
type AttachPortForwardingRuleParamDetail struct {
	RuleUuid string `json:"ruleUuid" validate:"required"`
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
}

// AttachPortForwardingRuleParam AttachPortForwardingRule request param
type AttachPortForwardingRuleParam struct {
	BaseParam
	AttachPortForwardingRule AttachPortForwardingRuleParamDetail `json:"attachPortForwardingRule"`
}
// CreatePortForwardingRuleParamDetail CreatePortForwardingRule detail param
type CreatePortForwardingRuleParamDetail struct {
	VipUuid string `json:"vipUuid" validate:"required"`
	VipPortStart int `json:"vipPortStart" validate:"required"`
	VipPortEnd int `json:"vipPortEnd,omitempty"`
	PrivatePortStart int `json:"privatePortStart,omitempty"`
	PrivatePortEnd int `json:"privatePortEnd,omitempty"`
	ProtocolType string `json:"protocolType" validate:"required"`
	VmNicUuid string `json:"vmNicUuid,omitempty"`
	AllowedCidr string `json:"allowedCidr,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePortForwardingRuleParam CreatePortForwardingRule request param
type CreatePortForwardingRuleParam struct {
	BaseParam
	CreatePortForwardingRule CreatePortForwardingRuleParamDetail `json:"createPortForwardingRule"`
}
