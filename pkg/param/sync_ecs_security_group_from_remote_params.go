// Copyright (c) ZStack.io, Inc.

package param

// SyncEcsSecurityGroupFromRemoteDetailParam SyncEcsSecurityGroupFromRemote详细参数
type SyncEcsSecurityGroupFromRemoteDetailParam struct {
	rest string `json:"ecsVpcUuid" validate:"required"` // 必填
	rest string `json:"securityGroupId,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncEcsSecurityGroupFromRemoteParam SyncEcsSecurityGroupFromRemote请求参数
type SyncEcsSecurityGroupFromRemoteParam struct {
	BaseParam
	Params SyncEcsSecurityGroupFromRemoteDetailParam `json:"params"` // 详细参数
}

