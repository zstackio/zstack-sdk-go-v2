// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsVpcRemoteDetailParam DeleteEcsVpcRemote detail param
type DeleteEcsVpcRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEcsVpcRemoteParam DeleteEcsVpcRemote request param
type DeleteEcsVpcRemoteParam struct {
	BaseParam
	Params DeleteEcsVpcRemoteDetailParam `json:"params"`
}
