// Copyright (c) ZStack.io, Inc.

package param

// CreateOssBackupBucketRemoteDetailParam CreateOssBackupBucketRemote detail param
type CreateOssBackupBucketRemoteDetailParam struct {
	RegionId string `json:"regionId" validate:"required"`
	OssDomain string `json:"ossDomain,omitempty"`
	OssKey string `json:"ossKey,omitempty"`
	OssSecret string `json:"ossSecret,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateOssBackupBucketRemoteParam CreateOssBackupBucketRemote request param
type CreateOssBackupBucketRemoteParam struct {
	BaseParam
	Params CreateOssBackupBucketRemoteDetailParam `json:"params"`
}
