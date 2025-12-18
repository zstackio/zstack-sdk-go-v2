// Copyright (c) ZStack.io, Inc.

package param

// CreatePortForwardingRuleDetailParam CreatePortForwardingRule detail param
type CreatePortForwardingRuleDetailParam struct {
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
	Params CreatePortForwardingRuleDetailParam `json:"params"`
}
