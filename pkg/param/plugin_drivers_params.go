// Copyright (c) ZStack.io, Inc.

package param

// DeletePluginDriversDetailParam DeletePluginDrivers详细参数
type DeletePluginDriversDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeletePluginDriversParam DeletePluginDrivers请求参数
type DeletePluginDriversParam struct {
	BaseParam
	Params DeletePluginDriversDetailParam `json:"params"` // 详细参数
}

