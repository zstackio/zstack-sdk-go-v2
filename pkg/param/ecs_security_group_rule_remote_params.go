// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsSecurityGroupRuleRemoteDetailParam DeleteEcsSecurityGroupRuleRemote详细参数
type DeleteEcsSecurityGroupRuleRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteEcsSecurityGroupRuleRemoteParam DeleteEcsSecurityGroupRuleRemote请求参数
type DeleteEcsSecurityGroupRuleRemoteParam struct {
	BaseParam
	Params DeleteEcsSecurityGroupRuleRemoteDetailParam `json:"params"` // 详细参数
}

