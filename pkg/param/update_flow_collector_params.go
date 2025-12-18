// Copyright (c) ZStack.io, Inc.

package param

// UpdateFlowCollectorDetailParam UpdateFlowCollector detail param
type UpdateFlowCollectorDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Server string `json:"server,omitempty"`
	Port int64 `json:"port,omitempty"`
}

// UpdateFlowCollectorParam UpdateFlowCollector request param
type UpdateFlowCollectorParam struct {
	BaseParam
	Params UpdateFlowCollectorDetailParam `json:"params"`
}
