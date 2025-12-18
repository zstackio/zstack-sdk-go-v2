// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsSecurityGroupInLocalDetailParam DeleteEcsSecurityGroupInLocal详细参数
type DeleteEcsSecurityGroupInLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteEcsSecurityGroupInLocalParam DeleteEcsSecurityGroupInLocal请求参数
type DeleteEcsSecurityGroupInLocalParam struct {
	BaseParam
	Params DeleteEcsSecurityGroupInLocalDetailParam `json:"params"` // 详细参数
}

