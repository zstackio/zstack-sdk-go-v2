// Copyright (c) ZStack.io, Inc.

package param

// RefreshPluginDriversDetailParam RefreshPluginDrivers详细参数
type RefreshPluginDriversDetailParam struct {
	rest string `json:"name,omitempty"`
}

// RefreshPluginDriversParam RefreshPluginDrivers请求参数
type RefreshPluginDriversParam struct {
	BaseParam
	Params RefreshPluginDriversDetailParam `json:"params"` // 详细参数
}

