// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsVSwitchRemoteDetailParam DeleteEcsVSwitchRemote详细参数
type DeleteEcsVSwitchRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteEcsVSwitchRemoteParam DeleteEcsVSwitchRemote请求参数
type DeleteEcsVSwitchRemoteParam struct {
	BaseParam
	Params DeleteEcsVSwitchRemoteDetailParam `json:"params"` // 详细参数
}

