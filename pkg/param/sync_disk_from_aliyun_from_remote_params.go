// Copyright (c) ZStack.io, Inc.

package param

// SyncDiskFromAliyunFromRemoteDetailParam SyncDiskFromAliyunFromRemote详细参数
type SyncDiskFromAliyunFromRemoteDetailParam struct {
	rest string `json:"identityUuid" validate:"required"` // 必填
	rest string `json:"diskId,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncDiskFromAliyunFromRemoteParam SyncDiskFromAliyunFromRemote请求参数
type SyncDiskFromAliyunFromRemoteParam struct {
	BaseParam
	Params SyncDiskFromAliyunFromRemoteDetailParam `json:"params"` // 详细参数
}

