// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// StartEcsInstanceParamDetail StartEcsInstance detail param
type StartEcsInstanceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// StartEcsInstanceParam StartEcsInstance request param
type StartEcsInstanceParam struct {
	BaseParam
	StartEcsInstance StartEcsInstanceParamDetail `json:"startEcsInstance"`
}
// DeleteEcsInstanceParamDetail DeleteEcsInstance detail param
type DeleteEcsInstanceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEcsInstanceParam DeleteEcsInstance request param
type DeleteEcsInstanceParam struct {
	BaseParam
	DeleteEcsInstance DeleteEcsInstanceParamDetail `json:"deleteEcsInstance"`
}
// StopEcsInstanceParamDetail StopEcsInstance detail param
type StopEcsInstanceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// StopEcsInstanceParam StopEcsInstance request param
type StopEcsInstanceParam struct {
	BaseParam
	StopEcsInstance StopEcsInstanceParamDetail `json:"stopEcsInstance"`
}
// RebootEcsInstanceParamDetail RebootEcsInstance detail param
type RebootEcsInstanceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RebootEcsInstanceParam RebootEcsInstance request param
type RebootEcsInstanceParam struct {
	BaseParam
	RebootEcsInstance RebootEcsInstanceParamDetail `json:"rebootEcsInstance"`
}
// UpdateEcsInstanceParamDetail UpdateEcsInstance detail param
type UpdateEcsInstanceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Password string `json:"password,omitempty"`
}

// UpdateEcsInstanceParam UpdateEcsInstance request param
type UpdateEcsInstanceParam struct {
	BaseParam
	UpdateEcsInstance UpdateEcsInstanceParamDetail `json:"updateEcsInstance"`
}
