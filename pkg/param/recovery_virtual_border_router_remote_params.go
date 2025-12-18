// Copyright (c) ZStack.io, Inc.

package param

// RecoveryVirtualBorderRouterRemoteDetailParam RecoveryVirtualBorderRouterRemote detail param
type RecoveryVirtualBorderRouterRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RecoveryVirtualBorderRouterRemoteParam RecoveryVirtualBorderRouterRemote request param
type RecoveryVirtualBorderRouterRemoteParam struct {
	BaseParam
	Params RecoveryVirtualBorderRouterRemoteDetailParam `json:"params"`
}
