// Copyright (c) ZStack.io, Inc.

package param

// RemoveRendezvousPointFromMulticastRouterDetailParam RemoveRendezvousPointFromMulticastRouter详细参数
type RemoveRendezvousPointFromMulticastRouterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"rpAddress" validate:"required"` // 必填
	rest string `json:"groupAddress" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// RemoveRendezvousPointFromMulticastRouterParam RemoveRendezvousPointFromMulticastRouter请求参数
type RemoveRendezvousPointFromMulticastRouterParam struct {
	BaseParam
	Params RemoveRendezvousPointFromMulticastRouterDetailParam `json:"params"` // 详细参数
}

