// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// ReconnectHostParamDetail ReconnectHost detail param
type ReconnectHostParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectHostParam ReconnectHost request param
type ReconnectHostParam struct {
	BaseParam
	Params ReconnectHostParamDetail `json:"params"`
}
// UpdateHostParamDetail UpdateHost detail param
type UpdateHostParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
}

// UpdateHostParam UpdateHost request param
type UpdateHostParam struct {
	BaseParam
	Params UpdateHostParamDetail `json:"params"`
}
// DeleteHostParamDetail DeleteHost detail param
type DeleteHostParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteHostParam DeleteHost request param
type DeleteHostParam struct {
	BaseParam
	Params DeleteHostParamDetail `json:"params"`
}
