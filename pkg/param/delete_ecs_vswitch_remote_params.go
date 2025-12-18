// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsVSwitchRemoteDetailParam DeleteEcsVSwitchRemote detail param
type DeleteEcsVSwitchRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEcsVSwitchRemoteParam DeleteEcsVSwitchRemote request param
type DeleteEcsVSwitchRemoteParam struct {
	BaseParam
	Params DeleteEcsVSwitchRemoteDetailParam `json:"params"`
}
