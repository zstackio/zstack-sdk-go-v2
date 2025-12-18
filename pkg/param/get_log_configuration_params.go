// Copyright (c) ZStack.io, Inc.

package param

// GetLogConfigurationDetailParam GetLogConfiguration detail param
type GetLogConfigurationDetailParam struct {
}

// GetLogConfigurationParam GetLogConfiguration request param
type GetLogConfigurationParam struct {
	BaseParam
	Params GetLogConfigurationDetailParam `json:"params"`
}
