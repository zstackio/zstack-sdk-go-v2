// Copyright (c) ZStack.io, Inc.

package param

// GetFlowMeterRouterIdDetailParam GetFlowMeterRouterId detail param
type GetFlowMeterRouterIdDetailParam struct {
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
}

// GetFlowMeterRouterIdParam GetFlowMeterRouterId request param
type GetFlowMeterRouterIdParam struct {
	BaseParam
	Params GetFlowMeterRouterIdDetailParam `json:"params"`
}
