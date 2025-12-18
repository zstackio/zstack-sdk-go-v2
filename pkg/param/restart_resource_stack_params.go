// Copyright (c) ZStack.io, Inc.

package param

// RestartResourceStackDetailParam RestartResourceStack详细参数
type RestartResourceStackDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// RestartResourceStackParam RestartResourceStack请求参数
type RestartResourceStackParam struct {
	BaseParam
	Params RestartResourceStackDetailParam `json:"params"` // 详细参数
}

