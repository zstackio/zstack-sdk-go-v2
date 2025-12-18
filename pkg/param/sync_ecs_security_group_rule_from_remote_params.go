// Copyright (c) ZStack.io, Inc.

package param

// SyncEcsSecurityGroupRuleFromRemoteDetailParam SyncEcsSecurityGroupRuleFromRemote详细参数
type SyncEcsSecurityGroupRuleFromRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncEcsSecurityGroupRuleFromRemoteParam SyncEcsSecurityGroupRuleFromRemote请求参数
type SyncEcsSecurityGroupRuleFromRemoteParam struct {
	BaseParam
	Params SyncEcsSecurityGroupRuleFromRemoteDetailParam `json:"params"` // 详细参数
}

