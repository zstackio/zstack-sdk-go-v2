// Copyright (c) ZStack.io, Inc.

package param

// SyncDataCenterFromRemoteDetailParam SyncDataCenterFromRemote详细参数
type SyncDataCenterFromRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncDataCenterFromRemoteParam SyncDataCenterFromRemote请求参数
type SyncDataCenterFromRemoteParam struct {
	BaseParam
	Params SyncDataCenterFromRemoteDetailParam `json:"params"` // 详细参数
}

