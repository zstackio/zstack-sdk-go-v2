// Copyright (c) ZStack.io, Inc.

package param

// SyncImageDetailParam SyncImage detail param
type SyncImageDetailParam struct {
	ImageStoreUuid string `json:"imageStoreUuid" validate:"required"`
}

// SyncImageParam SyncImage request param
type SyncImageParam struct {
	BaseParam
	Params SyncImageDetailParam `json:"params"`
}
