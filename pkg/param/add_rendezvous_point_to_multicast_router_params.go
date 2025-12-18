// Copyright (c) ZStack.io, Inc.

package param

// AddRendezvousPointToMulticastRouterDetailParam AddRendezvousPointToMulticastRouter detail param
type AddRendezvousPointToMulticastRouterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	RpAddress string `json:"rpAddress" validate:"required"`
	GroupAddress string `json:"groupAddress" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddRendezvousPointToMulticastRouterParam AddRendezvousPointToMulticastRouter request param
type AddRendezvousPointToMulticastRouterParam struct {
	BaseParam
	Params AddRendezvousPointToMulticastRouterDetailParam `json:"params"`
}
