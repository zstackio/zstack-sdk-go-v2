// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteL2NetworkParamDetail DeleteL2Network detail param
type DeleteL2NetworkParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteL2NetworkParam DeleteL2Network request param
type DeleteL2NetworkParam struct {
	BaseParam
	DeleteL2Network DeleteL2NetworkParamDetail `json:"deleteL2Network"`
}
// UpdateL2NetworkParamDetail UpdateL2Network detail param
type UpdateL2NetworkParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateL2NetworkParam UpdateL2Network request param
type UpdateL2NetworkParam struct {
	BaseParam
	UpdateL2Network UpdateL2NetworkParamDetail `json:"updateL2Network"`
}
