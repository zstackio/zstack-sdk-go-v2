// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateLogServerParamDetail UpdateLogServer detail param
type UpdateLogServerParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateLogServerParam UpdateLogServer request param
type UpdateLogServerParam struct {
	BaseParam
	Params UpdateLogServerParamDetail `json:"updateLogServer"`
}
// DeleteLogServerParamDetail DeleteLogServer detail param
type DeleteLogServerParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteLogServerParam DeleteLogServer request param
type DeleteLogServerParam struct {
	BaseParam
	Params DeleteLogServerParamDetail `json:"deleteLogServer"`
}
// AddLogServerParamDetail AddLogServer detail param
type AddLogServerParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Category string `json:"category" validate:"required"`
	Type string `json:"type" validate:"required"`
	Level string `json:"level,omitempty"`
	Configuration string `json:"configuration" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddLogServerParam AddLogServer request param
type AddLogServerParam struct {
	BaseParam
	Params AddLogServerParamDetail `json:"addLogServer"`
}
