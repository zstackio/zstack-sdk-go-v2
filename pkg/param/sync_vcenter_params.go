// Copyright (c) ZStack.io, Inc.

package param

// SyncVCenterDetailParam SyncVCenter详细参数
type SyncVCenterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// SyncVCenterParam SyncVCenter请求参数
type SyncVCenterParam struct {
	BaseParam
	Params SyncVCenterDetailParam `json:"params"` // 详细参数
}

