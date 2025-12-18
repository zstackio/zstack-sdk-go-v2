// Copyright (c) ZStack.io, Inc.

package param

// UpdateLogConfigurationDetailParam UpdateLogConfiguration detail param
type UpdateLogConfigurationDetailParam struct {
	ConfigId int64 `json:"configId" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateLogConfigurationParam UpdateLogConfiguration request param
type UpdateLogConfigurationParam struct {
	BaseParam
	Params UpdateLogConfigurationDetailParam `json:"params"`
}
