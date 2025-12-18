// Copyright (c) ZStack.io, Inc.

package param

// DeleteSSOClientDetailParam DeleteSSOClient detail param
type DeleteSSOClientDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSSOClientParam DeleteSSOClient request param
type DeleteSSOClientParam struct {
	BaseParam
	Params DeleteSSOClientDetailParam `json:"params"`
}
