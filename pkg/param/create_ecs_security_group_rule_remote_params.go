// Copyright (c) ZStack.io, Inc.

package param

// CreateEcsSecurityGroupRuleRemoteDetailParam CreateEcsSecurityGroupRuleRemote detail param
type CreateEcsSecurityGroupRuleRemoteDetailParam struct {
	GroupUuid string `json:"groupUuid" validate:"required"`
	Direction string `json:"direction" validate:"required"`
	Protocol string `json:"protocol" validate:"required"`
	PortRange string `json:"portRange" validate:"required"`
	Cidr string `json:"cidr" validate:"required"`
	Policy string `json:"policy,omitempty"`
	Nictype string `json:"nictype,omitempty"`
	Priority int `json:"priority,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEcsSecurityGroupRuleRemoteParam CreateEcsSecurityGroupRuleRemote request param
type CreateEcsSecurityGroupRuleRemoteParam struct {
	BaseParam
	Params CreateEcsSecurityGroupRuleRemoteDetailParam `json:"params"`
}
