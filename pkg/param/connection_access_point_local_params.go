// Copyright (c) ZStack.io, Inc.

package param

// DeleteConnectionAccessPointLocalDetailParam DeleteConnectionAccessPointLocal详细参数
type DeleteConnectionAccessPointLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteConnectionAccessPointLocalParam DeleteConnectionAccessPointLocal请求参数
type DeleteConnectionAccessPointLocalParam struct {
	BaseParam
	Params DeleteConnectionAccessPointLocalDetailParam `json:"params"` // 详细参数
}

