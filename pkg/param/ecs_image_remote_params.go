// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsImageRemoteDetailParam DeleteEcsImageRemote详细参数
type DeleteEcsImageRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteEcsImageRemoteParam DeleteEcsImageRemote请求参数
type DeleteEcsImageRemoteParam struct {
	BaseParam
	Params DeleteEcsImageRemoteDetailParam `json:"params"` // 详细参数
}

