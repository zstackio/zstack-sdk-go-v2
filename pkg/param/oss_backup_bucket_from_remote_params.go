// Copyright (c) ZStack.io, Inc.

package param

// GetOssBackupBucketFromRemoteDetailParam GetOssBackupBucketFromRemote详细参数
type GetOssBackupBucketFromRemoteDetailParam struct {
}

// GetOssBackupBucketFromRemoteParam GetOssBackupBucketFromRemote请求参数
type GetOssBackupBucketFromRemoteParam struct {
	BaseParam
	Params GetOssBackupBucketFromRemoteDetailParam `json:"params"` // 详细参数
}

