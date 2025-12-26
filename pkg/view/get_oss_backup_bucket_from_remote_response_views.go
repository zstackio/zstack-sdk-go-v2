// Copyright (c) ZStack.io, Inc.

package view

// GetOssBackupBucketFromRemoteView GetOssBackupBucketFromRemote
type GetOssBackupBucketFromRemoteView struct {
	Buckets []OssBucketFilesPropertyView `json:"buckets,omitempty"`
	Success bool `json:"success,omitempty"`
}

