// Copyright (c) ZStack.io, Inc.

package param

// DeleteOssBucketFileRemoteDetailParam DeleteOssBucketFileRemote detail param
type DeleteOssBucketFileRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	FileName string `json:"fileName" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteOssBucketFileRemoteParam DeleteOssBucketFileRemote request param
type DeleteOssBucketFileRemoteParam struct {
	BaseParam
	Params DeleteOssBucketFileRemoteDetailParam `json:"params"`
}
