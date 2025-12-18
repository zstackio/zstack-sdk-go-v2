// Copyright (c) ZStack.io, Inc.

package param

// SetFlowMeterRouterIdDetailParam SetFlowMeterRouterId detail param
type SetFlowMeterRouterIdDetailParam struct {
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
	RouterId int64 `json:"routerId" validate:"required"`
}

// SetFlowMeterRouterIdParam SetFlowMeterRouterId request param
type SetFlowMeterRouterIdParam struct {
	BaseParam
	Params SetFlowMeterRouterIdDetailParam `json:"params"`
}
