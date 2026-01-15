// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateFlowMeterParamDetail CreateFlowMeter detail param
type CreateFlowMeterParamDetail struct {
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
	CreateFlowMeter CreateFlowMeterParamDetail `json:"createFlowMeter"`
}
// DeleteFlowMeterParamDetail DeleteFlowMeter detail param
type DeleteFlowMeterParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteFlowMeterParam DeleteFlowMeter request param
type DeleteFlowMeterParam struct {
	BaseParam
	DeleteFlowMeter DeleteFlowMeterParamDetail `json:"deleteFlowMeter"`
}
// UpdateFlowMeterParamDetail UpdateFlowMeter detail param
type UpdateFlowMeterParamDetail struct {
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
	UpdateFlowMeter UpdateFlowMeterParamDetail `json:"updateFlowMeter"`
}
