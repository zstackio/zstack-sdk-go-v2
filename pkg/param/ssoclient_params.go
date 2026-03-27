// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteSSOClientParamDetail DeleteSSOClient detail param
type DeleteSSOClientParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteSSOClientParam DeleteSSOClient request param
type DeleteSSOClientParam struct {
	BaseParam
	Params DeleteSSOClientParamDetail `json:"params"`
}
// GetSSOClientParamDetail GetSSOClient detail param
type GetSSOClientParamDetail struct {
}

// GetSSOClientParam GetSSOClient request param
type GetSSOClientParam struct {
	BaseParam
	Params GetSSOClientParamDetail `json:"getSSOClient"`
}
