// Copyright (c) ZStack.io, Inc.

package param

// SyncVmClockDetailParam SyncVmClock详细参数
type SyncVmClockDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// SyncVmClockParam SyncVmClock请求参数
type SyncVmClockParam struct {
	BaseParam
	Params SyncVmClockDetailParam `json:"params"` // 详细参数
}

