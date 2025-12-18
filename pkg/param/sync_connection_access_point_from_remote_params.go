// Copyright (c) ZStack.io, Inc.

package param

// SyncConnectionAccessPointFromRemoteDetailParam SyncConnectionAccessPointFromRemote详细参数
type SyncConnectionAccessPointFromRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"accessPointId,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncConnectionAccessPointFromRemoteParam SyncConnectionAccessPointFromRemote请求参数
type SyncConnectionAccessPointFromRemoteParam struct {
	BaseParam
	Params SyncConnectionAccessPointFromRemoteDetailParam `json:"params"` // 详细参数
}

