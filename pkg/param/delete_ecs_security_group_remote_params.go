// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsSecurityGroupRemoteDetailParam DeleteEcsSecurityGroupRemote detail param
type DeleteEcsSecurityGroupRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEcsSecurityGroupRemoteParam DeleteEcsSecurityGroupRemote request param
type DeleteEcsSecurityGroupRemoteParam struct {
	BaseParam
	Params DeleteEcsSecurityGroupRemoteDetailParam `json:"params"`
}
