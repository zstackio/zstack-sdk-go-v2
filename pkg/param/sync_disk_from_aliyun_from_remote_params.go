// Copyright (c) ZStack.io, Inc.

package param

// SyncDiskFromAliyunFromRemoteDetailParam SyncDiskFromAliyunFromRemote detail param
type SyncDiskFromAliyunFromRemoteDetailParam struct {
	IdentityUuid string `json:"identityUuid" validate:"required"`
	DiskId string `json:"diskId,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncDiskFromAliyunFromRemoteParam SyncDiskFromAliyunFromRemote request param
type SyncDiskFromAliyunFromRemoteParam struct {
	BaseParam
	Params SyncDiskFromAliyunFromRemoteDetailParam `json:"params"`
}
