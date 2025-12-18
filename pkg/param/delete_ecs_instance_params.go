// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsInstanceDetailParam DeleteEcsInstance detail param
type DeleteEcsInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEcsInstanceParam DeleteEcsInstance request param
type DeleteEcsInstanceParam struct {
	BaseParam
	Params DeleteEcsInstanceDetailParam `json:"params"`
}
