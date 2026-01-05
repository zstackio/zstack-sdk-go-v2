// Copyright (c) ZStack.io, Inc.

package param

// SyncNfvInstGroupDetailParam SyncNfvInstGroup detail param
type SyncNfvInstGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// SyncNfvInstGroupParam SyncNfvInstGroup request param
type SyncNfvInstGroupParam struct {
	BaseParam
	Params SyncNfvInstGroupDetailParam `json:"params"`
}
