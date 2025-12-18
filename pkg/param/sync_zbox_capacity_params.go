// Copyright (c) ZStack.io, Inc.

package param

// SyncZBoxCapacityDetailParam SyncZBoxCapacity详细参数
type SyncZBoxCapacityDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// SyncZBoxCapacityParam SyncZBoxCapacity请求参数
type SyncZBoxCapacityParam struct {
	BaseParam
	Params SyncZBoxCapacityDetailParam `json:"params"` // 详细参数
}

