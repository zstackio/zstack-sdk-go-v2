// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsVSwitchInLocalDetailParam DeleteEcsVSwitchInLocal detail param
type DeleteEcsVSwitchInLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEcsVSwitchInLocalParam DeleteEcsVSwitchInLocal request param
type DeleteEcsVSwitchInLocalParam struct {
	BaseParam
	Params DeleteEcsVSwitchInLocalDetailParam `json:"params"`
}
