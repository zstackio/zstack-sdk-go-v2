// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsImageRemoteDetailParam DeleteEcsImageRemote detail param
type DeleteEcsImageRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEcsImageRemoteParam DeleteEcsImageRemote request param
type DeleteEcsImageRemoteParam struct {
	BaseParam
	Params DeleteEcsImageRemoteDetailParam `json:"params"`
}
