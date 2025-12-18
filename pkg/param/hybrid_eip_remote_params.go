// Copyright (c) ZStack.io, Inc.

package param

// DeleteHybridEipRemoteDetailParam DeleteHybridEipRemote详细参数
type DeleteHybridEipRemoteDetailParam struct {
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteHybridEipRemoteParam DeleteHybridEipRemote请求参数
type DeleteHybridEipRemoteParam struct {
	BaseParam
	Params DeleteHybridEipRemoteDetailParam `json:"params"` // 详细参数
}

