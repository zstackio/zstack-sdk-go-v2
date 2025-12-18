// Copyright (c) ZStack.io, Inc.

package param

// ChangeZoneStateDetailParam ChangeZoneState详细参数
type ChangeZoneStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeZoneStateParam ChangeZoneState请求参数
type ChangeZoneStateParam struct {
	BaseParam
	Params ChangeZoneStateDetailParam `json:"params"` // 详细参数
}

