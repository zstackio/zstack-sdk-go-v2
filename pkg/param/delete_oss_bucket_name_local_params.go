// Copyright (c) ZStack.io, Inc.

package param

// DeleteOssBucketNameLocalDetailParam DeleteOssBucketNameLocal detail param
type DeleteOssBucketNameLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteOssBucketNameLocalParam DeleteOssBucketNameLocal request param
type DeleteOssBucketNameLocalParam struct {
	BaseParam
	Params DeleteOssBucketNameLocalDetailParam `json:"params"`
}
