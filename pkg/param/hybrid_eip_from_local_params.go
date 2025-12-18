// Copyright (c) ZStack.io, Inc.

package param

// DeleteHybridEipFromLocalDetailParam DeleteHybridEipFromLocal详细参数
type DeleteHybridEipFromLocalDetailParam struct {
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteHybridEipFromLocalParam DeleteHybridEipFromLocal请求参数
type DeleteHybridEipFromLocalParam struct {
	BaseParam
	Params DeleteHybridEipFromLocalDetailParam `json:"params"` // 详细参数
}

