// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsImageLocalDetailParam DeleteEcsImageLocal详细参数
type DeleteEcsImageLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteEcsImageLocalParam DeleteEcsImageLocal请求参数
type DeleteEcsImageLocalParam struct {
	BaseParam
	Params DeleteEcsImageLocalDetailParam `json:"params"` // 详细参数
}

