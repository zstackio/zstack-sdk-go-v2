// Copyright (c) ZStack.io, Inc.

package param

// DeleteIdentityZoneInLocalDetailParam DeleteIdentityZoneInLocal详细参数
type DeleteIdentityZoneInLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteIdentityZoneInLocalParam DeleteIdentityZoneInLocal请求参数
type DeleteIdentityZoneInLocalParam struct {
	BaseParam
	Params DeleteIdentityZoneInLocalDetailParam `json:"params"` // 详细参数
}

