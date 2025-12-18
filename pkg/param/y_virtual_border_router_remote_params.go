// Copyright (c) ZStack.io, Inc.

package param

// RecoveryVirtualBorderRouterRemoteDetailParam RecoveryVirtualBorderRouterRemote详细参数
type RecoveryVirtualBorderRouterRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// RecoveryVirtualBorderRouterRemoteParam RecoveryVirtualBorderRouterRemote请求参数
type RecoveryVirtualBorderRouterRemoteParam struct {
	BaseParam
	Params RecoveryVirtualBorderRouterRemoteDetailParam `json:"params"` // 详细参数
}

