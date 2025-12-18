// Copyright (c) ZStack.io, Inc.

package param

// RemoveVRouterNetworksFromFlowMeterDetailParam RemoveVRouterNetworksFromFlowMeter详细参数
type RemoveVRouterNetworksFromFlowMeterDetailParam struct {
	rest []string `json:"uuids" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// RemoveVRouterNetworksFromFlowMeterParam RemoveVRouterNetworksFromFlowMeter请求参数
type RemoveVRouterNetworksFromFlowMeterParam struct {
	BaseParam
	Params RemoveVRouterNetworksFromFlowMeterDetailParam `json:"params"` // 详细参数
}

