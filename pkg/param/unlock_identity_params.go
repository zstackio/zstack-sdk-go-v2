// Copyright (c) ZStack.io, Inc.

package param

// UnlockIdentityDetailParam UnlockIdentity detail param
type UnlockIdentityDetailParam struct {
	ResourceName string `json:"resourceName" validate:"required"`
	LoginType string `json:"loginType" validate:"required"`
}

// UnlockIdentityParam UnlockIdentity request param
type UnlockIdentityParam struct {
	BaseParam
	Params UnlockIdentityDetailParam `json:"params"`
}
