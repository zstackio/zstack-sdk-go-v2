// Copyright (c) ZStack.io, Inc.

package param

// SyncVCenterDetailParam SyncVCenter detail param
type SyncVCenterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// SyncVCenterParam SyncVCenter request param
type SyncVCenterParam struct {
	BaseParam
	Params SyncVCenterDetailParam `json:"params"`
}
