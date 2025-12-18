// Copyright (c) ZStack.io, Inc.

package param

// RemoveVRouterNetworksFromOspfAreaDetailParam RemoveVRouterNetworksFromOspfArea详细参数
type RemoveVRouterNetworksFromOspfAreaDetailParam struct {
	rest []string `json:"uuids" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// RemoveVRouterNetworksFromOspfAreaParam RemoveVRouterNetworksFromOspfArea请求参数
type RemoveVRouterNetworksFromOspfAreaParam struct {
	BaseParam
	Params RemoveVRouterNetworksFromOspfAreaDetailParam `json:"params"` // 详细参数
}

