// Copyright (c) ZStack.io, Inc.

package param

// SyncImageSizeDetailParam SyncImageSize detail param
type SyncImageSizeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// SyncImageSizeParam SyncImageSize request param
type SyncImageSizeParam struct {
	BaseParam
	Params SyncImageSizeDetailParam `json:"params"`
}
