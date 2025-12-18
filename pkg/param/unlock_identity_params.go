// Copyright (c) ZStack.io, Inc.

package param

// UnlockIdentityDetailParam UnlockIdentity详细参数
type UnlockIdentityDetailParam struct {
	rest string `json:"resourceName" validate:"required"` // 必填
	rest string `json:"loginType" validate:"required"` // 必填
}

// UnlockIdentityParam UnlockIdentity请求参数
type UnlockIdentityParam struct {
	BaseParam
	Params UnlockIdentityDetailParam `json:"params"` // 详细参数
}

