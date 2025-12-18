// Copyright (c) ZStack.io, Inc.

package param

// GetVRouterFlowCounterDetailParam GetVRouterFlowCounter详细参数
type GetVRouterFlowCounterDetailParam struct {
	rest string `json:"vRouterUuid" validate:"required"` // 必填
}

// GetVRouterFlowCounterParam GetVRouterFlowCounter请求参数
type GetVRouterFlowCounterParam struct {
	BaseParam
	Params GetVRouterFlowCounterDetailParam `json:"params"` // 详细参数
}

