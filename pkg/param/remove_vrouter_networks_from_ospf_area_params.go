// Copyright (c) ZStack.io, Inc.

package param

// RemoveVRouterNetworksFromOspfAreaDetailParam RemoveVRouterNetworksFromOspfArea detail param
type RemoveVRouterNetworksFromOspfAreaDetailParam struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RemoveVRouterNetworksFromOspfAreaParam RemoveVRouterNetworksFromOspfArea request param
type RemoveVRouterNetworksFromOspfAreaParam struct {
	BaseParam
	Params RemoveVRouterNetworksFromOspfAreaDetailParam `json:"params"`
}
