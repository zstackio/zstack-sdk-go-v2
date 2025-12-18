// Copyright (c) ZStack.io, Inc.

package param

// SyncEcsVpcFromRemoteDetailParam SyncEcsVpcFromRemote详细参数
type SyncEcsVpcFromRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"ecsVpcId,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncEcsVpcFromRemoteParam SyncEcsVpcFromRemote请求参数
type SyncEcsVpcFromRemoteParam struct {
	BaseParam
	Params SyncEcsVpcFromRemoteDetailParam `json:"params"` // 详细参数
}

