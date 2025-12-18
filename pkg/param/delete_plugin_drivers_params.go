// Copyright (c) ZStack.io, Inc.

package param

// DeletePluginDriversDetailParam DeletePluginDrivers detail param
type DeletePluginDriversDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePluginDriversParam DeletePluginDrivers request param
type DeletePluginDriversParam struct {
	BaseParam
	Params DeletePluginDriversDetailParam `json:"params"`
}
