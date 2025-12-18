// Copyright (c) ZStack.io, Inc.

package param

// AttachNicToBondingDetailParam AttachNicToBonding详细参数
type AttachNicToBondingDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"slaveUuids" validate:"required"` // 必填
	rest string `json:"type,omitempty"`
}

// AttachNicToBondingParam AttachNicToBonding请求参数
type AttachNicToBondingParam struct {
	BaseParam
	Params AttachNicToBondingDetailParam `json:"params"` // 详细参数
}

