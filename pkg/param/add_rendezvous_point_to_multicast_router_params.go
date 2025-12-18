// Copyright (c) ZStack.io, Inc.

package param

// AddRendezvousPointToMulticastRouterDetailParam AddRendezvousPointToMulticastRouter详细参数
type AddRendezvousPointToMulticastRouterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"rpAddress" validate:"required"` // 必填
	rest string `json:"groupAddress" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddRendezvousPointToMulticastRouterParam AddRendezvousPointToMulticastRouter请求参数
type AddRendezvousPointToMulticastRouterParam struct {
	BaseParam
	Params AddRendezvousPointToMulticastRouterDetailParam `json:"params"` // 详细参数
}

