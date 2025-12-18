// Copyright (c) ZStack.io, Inc.

package param

// SyncImageDetailParam SyncImage详细参数
type SyncImageDetailParam struct {
	rest string `json:"imageStoreUuid" validate:"required"` // 必填
}

// SyncImageParam SyncImage请求参数
type SyncImageParam struct {
	BaseParam
	Params SyncImageDetailParam `json:"params"` // 详细参数
}

