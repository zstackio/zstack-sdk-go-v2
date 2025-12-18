// Copyright (c) ZStack.io, Inc.

package param

// DeleteFlowMeterDetailParam DeleteFlowMeter detail param
type DeleteFlowMeterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteFlowMeterParam DeleteFlowMeter request param
type DeleteFlowMeterParam struct {
	BaseParam
	Params DeleteFlowMeterDetailParam `json:"params"`
}
