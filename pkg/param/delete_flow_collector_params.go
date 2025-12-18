// Copyright (c) ZStack.io, Inc.

package param

// DeleteFlowCollectorDetailParam DeleteFlowCollector detail param
type DeleteFlowCollectorDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteFlowCollectorParam DeleteFlowCollector request param
type DeleteFlowCollectorParam struct {
	BaseParam
	Params DeleteFlowCollectorDetailParam `json:"params"`
}
