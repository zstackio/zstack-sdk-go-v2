// Copyright (c) ZStack.io, Inc.

package param

// SyncEcsSecurityGroupRuleFromRemoteDetailParam SyncEcsSecurityGroupRuleFromRemote detail param
type SyncEcsSecurityGroupRuleFromRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncEcsSecurityGroupRuleFromRemoteParam SyncEcsSecurityGroupRuleFromRemote request param
type SyncEcsSecurityGroupRuleFromRemoteParam struct {
	BaseParam
	Params SyncEcsSecurityGroupRuleFromRemoteDetailParam `json:"params"`
}
