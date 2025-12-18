// Copyright (c) ZStack.io, Inc.

package param

// DeleteAlertDetailParam DeleteAlert detail param
type DeleteAlertDetailParam struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAlertParam DeleteAlert request param
type DeleteAlertParam struct {
	BaseParam
	Params DeleteAlertDetailParam `json:"params"`
}
