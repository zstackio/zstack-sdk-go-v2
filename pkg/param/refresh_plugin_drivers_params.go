// Copyright (c) ZStack.io, Inc.

package param

// RefreshPluginDriversDetailParam RefreshPluginDrivers detail param
type RefreshPluginDriversDetailParam struct {
	Name string `json:"name,omitempty"`
}

// RefreshPluginDriversParam RefreshPluginDrivers request param
type RefreshPluginDriversParam struct {
	BaseParam
	Params RefreshPluginDriversDetailParam `json:"params"`
}
