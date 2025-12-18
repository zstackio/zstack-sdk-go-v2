// Copyright (c) ZStack.io, Inc.

package param

// GetOssBackupBucketFromRemoteDetailParam GetOssBackupBucketFromRemote detail param
type GetOssBackupBucketFromRemoteDetailParam struct {
}

// GetOssBackupBucketFromRemoteParam GetOssBackupBucketFromRemote request param
type GetOssBackupBucketFromRemoteParam struct {
	BaseParam
	Params GetOssBackupBucketFromRemoteDetailParam `json:"params"`
}
