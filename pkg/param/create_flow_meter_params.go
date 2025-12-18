// Copyright (c) ZStack.io, Inc.

package param

// CreateFlowMeterDetailParam CreateFlowMeter detail param
type CreateFlowMeterDetailParam struct {
	Version string `json:"version,omitempty"`
	Type string `json:"type" validate:"required"`
	Sample int `json:"sample,omitempty"`
	GenerateInterval int `json:"generateInterval,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Server string `json:"server,omitempty"`
	Port int64 `json:"port,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFlowMeterParam CreateFlowMeter request param
type CreateFlowMeterParam struct {
	BaseParam
	Params CreateFlowMeterDetailParam `json:"params"`
}
