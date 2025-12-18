// Copyright (c) ZStack.io, Inc.

package param

// GetFlowMeterRouterIdDetailParam GetFlowMeterRouterId详细参数
type GetFlowMeterRouterIdDetailParam struct {
	rest string `json:"vRouterUuid" validate:"required"` // 必填
}

// GetFlowMeterRouterIdParam GetFlowMeterRouterId请求参数
type GetFlowMeterRouterIdParam struct {
	BaseParam
	Params GetFlowMeterRouterIdDetailParam `json:"params"` // 详细参数
}

