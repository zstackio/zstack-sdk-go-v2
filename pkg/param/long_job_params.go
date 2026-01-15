// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CleanLongJobParamDetail CleanLongJob detail param
type CleanLongJobParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// CleanLongJobParam CleanLongJob request param
type CleanLongJobParam struct {
	BaseParam
	CleanLongJob CleanLongJobParamDetail `json:"cleanLongJob"`
}
// ResumeLongJobParamDetail ResumeLongJob detail param
type ResumeLongJobParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ResumeLongJobParam ResumeLongJob request param
type ResumeLongJobParam struct {
	BaseParam
	ResumeLongJob ResumeLongJobParamDetail `json:"resumeLongJob"`
}
// DeleteLongJobParamDetail DeleteLongJob detail param
type DeleteLongJobParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteLongJobParam DeleteLongJob request param
type DeleteLongJobParam struct {
	BaseParam
	DeleteLongJob DeleteLongJobParamDetail `json:"deleteLongJob"`
}
// UpdateLongJobParamDetail UpdateLongJob detail param
type UpdateLongJobParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateLongJobParam UpdateLongJob request param
type UpdateLongJobParam struct {
	BaseParam
	UpdateLongJob UpdateLongJobParamDetail `json:"updateLongJob"`
}
