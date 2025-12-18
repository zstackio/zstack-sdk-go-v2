// Copyright (c) ZStack.io, Inc.

package param

// SyncEcsInstanceFromRemoteDetailParam SyncEcsInstanceFromRemote详细参数
type SyncEcsInstanceFromRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest bool `json:"onlyZstack,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncEcsInstanceFromRemoteParam SyncEcsInstanceFromRemote请求参数
type SyncEcsInstanceFromRemoteParam struct {
	BaseParam
	Params SyncEcsInstanceFromRemoteDetailParam `json:"params"` // 详细参数
}

