// Copyright (c) ZStack.io, Inc.

package param

// SyncIdentityFromRemoteDetailParam SyncIdentityFromRemote详细参数
type SyncIdentityFromRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncIdentityFromRemoteParam SyncIdentityFromRemote请求参数
type SyncIdentityFromRemoteParam struct {
	BaseParam
	Params SyncIdentityFromRemoteDetailParam `json:"params"` // 详细参数
}

