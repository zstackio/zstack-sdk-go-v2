// Copyright (c) ZStack.io, Inc.

package param

// TerminateVirtualBorderRouterRemoteDetailParam TerminateVirtualBorderRouterRemote detail param
type TerminateVirtualBorderRouterRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// TerminateVirtualBorderRouterRemoteParam TerminateVirtualBorderRouterRemote request param
type TerminateVirtualBorderRouterRemoteParam struct {
	BaseParam
	Params TerminateVirtualBorderRouterRemoteDetailParam `json:"params"`
}
