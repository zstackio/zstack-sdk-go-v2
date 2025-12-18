// Copyright (c) ZStack.io, Inc.

package param

// AddVRouterNetworksToFlowMeterDetailParam AddVRouterNetworksToFlowMeter详细参数
type AddVRouterNetworksToFlowMeterDetailParam struct {
	rest string `json:"flowMeterUuid" validate:"required"` // 必填
	rest string `json:"vRouterUuid" validate:"required"` // 必填
	rest []string `json:"l3NetworkUuids" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddVRouterNetworksToFlowMeterParam AddVRouterNetworksToFlowMeter请求参数
type AddVRouterNetworksToFlowMeterParam struct {
	BaseParam
	Params AddVRouterNetworksToFlowMeterDetailParam `json:"params"` // 详细参数
}

