// Copyright (c) ZStack.io, Inc.

package param

// SyncEcsSecurityGroupFromRemoteDetailParam SyncEcsSecurityGroupFromRemote detail param
type SyncEcsSecurityGroupFromRemoteDetailParam struct {
	EcsVpcUuid string `json:"ecsVpcUuid" validate:"required"`
	SecurityGroupId string `json:"securityGroupId,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncEcsSecurityGroupFromRemoteParam SyncEcsSecurityGroupFromRemote request param
type SyncEcsSecurityGroupFromRemoteParam struct {
	BaseParam
	Params SyncEcsSecurityGroupFromRemoteDetailParam `json:"params"`
}
