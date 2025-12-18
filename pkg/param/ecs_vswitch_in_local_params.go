// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsVSwitchInLocalDetailParam DeleteEcsVSwitchInLocal详细参数
type DeleteEcsVSwitchInLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteEcsVSwitchInLocalParam DeleteEcsVSwitchInLocal请求参数
type DeleteEcsVSwitchInLocalParam struct {
	BaseParam
	Params DeleteEcsVSwitchInLocalDetailParam `json:"params"` // 详细参数
}

