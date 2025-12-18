// Copyright (c) ZStack.io, Inc.

package param

// GetVRouterFlowCounterDetailParam GetVRouterFlowCounter detail param
type GetVRouterFlowCounterDetailParam struct {
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
}

// GetVRouterFlowCounterParam GetVRouterFlowCounter request param
type GetVRouterFlowCounterParam struct {
	BaseParam
	Params GetVRouterFlowCounterDetailParam `json:"params"`
}
