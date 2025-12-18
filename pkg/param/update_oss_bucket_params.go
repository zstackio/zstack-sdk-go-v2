// Copyright (c) ZStack.io, Inc.

package param

// UpdateOssBucketDetailParam UpdateOssBucket detail param
type UpdateOssBucketDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Description string `json:"description,omitempty"`
	OssDomain string `json:"ossDomain,omitempty"`
	OssKey string `json:"ossKey,omitempty"`
	OssSecret string `json:"ossSecret,omitempty"`
}

// UpdateOssBucketParam UpdateOssBucket request param
type UpdateOssBucketParam struct {
	BaseParam
	Params UpdateOssBucketDetailParam `json:"params"`
}
