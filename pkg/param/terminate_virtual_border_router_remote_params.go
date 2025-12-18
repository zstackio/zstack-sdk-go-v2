// Copyright (c) ZStack.io, Inc.

package param

// TerminateVirtualBorderRouterRemoteDetailParam TerminateVirtualBorderRouterRemote详细参数
type TerminateVirtualBorderRouterRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// TerminateVirtualBorderRouterRemoteParam TerminateVirtualBorderRouterRemote请求参数
type TerminateVirtualBorderRouterRemoteParam struct {
	BaseParam
	Params TerminateVirtualBorderRouterRemoteDetailParam `json:"params"` // 详细参数
}

