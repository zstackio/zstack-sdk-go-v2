// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsSecurityGroupRemoteDetailParam DeleteEcsSecurityGroupRemote详细参数
type DeleteEcsSecurityGroupRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteEcsSecurityGroupRemoteParam DeleteEcsSecurityGroupRemote请求参数
type DeleteEcsSecurityGroupRemoteParam struct {
	BaseParam
	Params DeleteEcsSecurityGroupRemoteDetailParam `json:"params"` // 详细参数
}

