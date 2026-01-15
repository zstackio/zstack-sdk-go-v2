// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateFlowCollectorParamDetail CreateFlowCollector detail param
type CreateFlowCollectorParamDetail struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	FlowMeterUuid string `json:"flowMeterUuid" validate:"required"`
	Server string `json:"server,omitempty"`
	Port int64 `json:"port,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFlowCollectorParam CreateFlowCollector request param
type CreateFlowCollectorParam struct {
	BaseParam
	CreateFlowCollector CreateFlowCollectorParamDetail `json:"createFlowCollector"`
}
// UpdateFlowCollectorParamDetail UpdateFlowCollector detail param
type UpdateFlowCollectorParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Server string `json:"server,omitempty"`
	Port int64 `json:"port,omitempty"`
}

// UpdateFlowCollectorParam UpdateFlowCollector request param
type UpdateFlowCollectorParam struct {
	BaseParam
	UpdateFlowCollector UpdateFlowCollectorParamDetail `json:"updateFlowCollector"`
}
// DeleteFlowCollectorParamDetail DeleteFlowCollector detail param
type DeleteFlowCollectorParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteFlowCollectorParam DeleteFlowCollector request param
type DeleteFlowCollectorParam struct {
	BaseParam
	DeleteFlowCollector DeleteFlowCollectorParamDetail `json:"deleteFlowCollector"`
}
