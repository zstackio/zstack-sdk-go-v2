// Copyright (c) ZStack.io, Inc.

package param

// SyncEcsImageFromRemoteDetailParam SyncEcsImageFromRemote详细参数
type SyncEcsImageFromRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"type,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncEcsImageFromRemoteParam SyncEcsImageFromRemote请求参数
type SyncEcsImageFromRemoteParam struct {
	BaseParam
	Params SyncEcsImageFromRemoteDetailParam `json:"params"` // 详细参数
}

