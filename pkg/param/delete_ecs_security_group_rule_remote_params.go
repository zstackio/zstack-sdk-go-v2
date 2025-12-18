// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsSecurityGroupRuleRemoteDetailParam DeleteEcsSecurityGroupRuleRemote detail param
type DeleteEcsSecurityGroupRuleRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEcsSecurityGroupRuleRemoteParam DeleteEcsSecurityGroupRuleRemote request param
type DeleteEcsSecurityGroupRuleRemoteParam struct {
	BaseParam
	Params DeleteEcsSecurityGroupRuleRemoteDetailParam `json:"params"`
}
