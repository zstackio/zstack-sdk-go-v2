// Copyright (c) ZStack.io, Inc.

package param

// SyncHybridEipFromRemoteDetailParam SyncHybridEipFromRemote详细参数
type SyncHybridEipFromRemoteDetailParam struct {
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncHybridEipFromRemoteParam SyncHybridEipFromRemote请求参数
type SyncHybridEipFromRemoteParam struct {
	BaseParam
	Params SyncHybridEipFromRemoteDetailParam `json:"params"` // 详细参数
}

