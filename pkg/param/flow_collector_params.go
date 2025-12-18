// Copyright (c) ZStack.io, Inc.

package param

// DeleteFlowCollectorDetailParam DeleteFlowCollector详细参数
type DeleteFlowCollectorDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteFlowCollectorParam DeleteFlowCollector请求参数
type DeleteFlowCollectorParam struct {
	BaseParam
	Params DeleteFlowCollectorDetailParam `json:"params"` // 详细参数
}

