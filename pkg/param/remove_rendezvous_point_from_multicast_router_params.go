// Copyright (c) ZStack.io, Inc.

package param

// RemoveRendezvousPointFromMulticastRouterDetailParam RemoveRendezvousPointFromMulticastRouter detail param
type RemoveRendezvousPointFromMulticastRouterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	RpAddress string `json:"rpAddress" validate:"required"`
	GroupAddress string `json:"groupAddress" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RemoveRendezvousPointFromMulticastRouterParam RemoveRendezvousPointFromMulticastRouter request param
type RemoveRendezvousPointFromMulticastRouterParam struct {
	BaseParam
	Params RemoveRendezvousPointFromMulticastRouterDetailParam `json:"params"`
}
