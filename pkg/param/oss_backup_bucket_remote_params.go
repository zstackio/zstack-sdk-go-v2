// Copyright (c) ZStack.io, Inc.

package param

// CreateOssBackupBucketRemoteDetailParam CreateOssBackupBucketRemote详细参数
type CreateOssBackupBucketRemoteDetailParam struct {
	rest string `json:"regionId" validate:"required"` // 必填
	rest string `json:"ossDomain,omitempty"`
	rest string `json:"ossKey,omitempty"`
	rest string `json:"ossSecret,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateOssBackupBucketRemoteParam CreateOssBackupBucketRemote请求参数
type CreateOssBackupBucketRemoteParam struct {
	BaseParam
	Params CreateOssBackupBucketRemoteDetailParam `json:"params"` // 详细参数
}

