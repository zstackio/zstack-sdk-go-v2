// Copyright (c) ZStack.io, Inc.

package param

// DeleteOssBucketRemoteDetailParam DeleteOssBucketRemote detail param
type DeleteOssBucketRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteOssBucketRemoteParam DeleteOssBucketRemote request param
type DeleteOssBucketRemoteParam struct {
	BaseParam
	Params DeleteOssBucketRemoteDetailParam `json:"params"`
}
