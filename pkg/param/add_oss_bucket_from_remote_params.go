// Copyright (c) ZStack.io, Inc.

package param

// AddOssBucketFromRemoteDetailParam AddOssBucketFromRemote detail param
type AddOssBucketFromRemoteDetailParam struct {
	BucketName string `json:"bucketName" validate:"required"`
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	Description string `json:"description,omitempty"`
	OssDomain string `json:"ossDomain,omitempty"`
	OssKey string `json:"ossKey,omitempty"`
	OssSecret string `json:"ossSecret,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddOssBucketFromRemoteParam AddOssBucketFromRemote request param
type AddOssBucketFromRemoteParam struct {
	BaseParam
	Params AddOssBucketFromRemoteDetailParam `json:"params"`
}
