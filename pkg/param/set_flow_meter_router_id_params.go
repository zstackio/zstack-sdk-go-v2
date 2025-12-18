// Copyright (c) ZStack.io, Inc.

package param

// SetFlowMeterRouterIdDetailParam SetFlowMeterRouterId详细参数
type SetFlowMeterRouterIdDetailParam struct {
	rest string `json:"vRouterUuid" validate:"required"` // 必填
	rest int64 `json:"routerId" validate:"required"` // 必填
}

// SetFlowMeterRouterIdParam SetFlowMeterRouterId请求参数
type SetFlowMeterRouterIdParam struct {
	BaseParam
	Params SetFlowMeterRouterIdDetailParam `json:"params"` // 详细参数
}

