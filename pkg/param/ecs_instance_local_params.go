// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsInstanceLocalDetailParam DeleteEcsInstanceLocal详细参数
type DeleteEcsInstanceLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteEcsInstanceLocalParam DeleteEcsInstanceLocal请求参数
type DeleteEcsInstanceLocalParam struct {
	BaseParam
	Params DeleteEcsInstanceLocalDetailParam `json:"params"` // 详细参数
}

