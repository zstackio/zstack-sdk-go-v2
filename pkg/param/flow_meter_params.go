// Copyright (c) ZStack.io, Inc.

package param

// DeleteFlowMeterDetailParam DeleteFlowMeter详细参数
type DeleteFlowMeterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteFlowMeterParam DeleteFlowMeter请求参数
type DeleteFlowMeterParam struct {
	BaseParam
	Params DeleteFlowMeterDetailParam `json:"params"` // 详细参数
}

