// Copyright (c) ZStack.io, Inc.

package param

// CreateOssBucketRemoteDetailParam CreateOssBucketRemote detail param
type CreateOssBucketRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	BucketName string `json:"bucketName" validate:"required"`
	Description string `json:"description,omitempty"`
	OssDomain string `json:"ossDomain,omitempty"`
	OssKey string `json:"ossKey,omitempty"`
	OssSecret string `json:"ossSecret,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateOssBucketRemoteParam CreateOssBucketRemote request param
type CreateOssBucketRemoteParam struct {
	BaseParam
	Params CreateOssBucketRemoteDetailParam `json:"params"`
}
