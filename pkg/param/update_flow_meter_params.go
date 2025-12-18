// Copyright (c) ZStack.io, Inc.

package param

// UpdateFlowMeterDetailParam UpdateFlowMeter detail param
type UpdateFlowMeterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Version string `json:"version,omitempty"`
	Sample int64 `json:"sample,omitempty"`
	Name string `json:"name,omitempty"`
	ExpireInterval int64 `json:"expireInterval,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateFlowMeterParam UpdateFlowMeter request param
type UpdateFlowMeterParam struct {
	BaseParam
	Params UpdateFlowMeterDetailParam `json:"params"`
}
