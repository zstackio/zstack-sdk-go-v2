// Copyright (c) ZStack.io, Inc.

package param

// DeleteLogConfigurationDetailParam DeleteLogConfiguration detail param
type DeleteLogConfigurationDetailParam struct {
	ConfigId int64 `json:"configId" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteLogConfigurationParam DeleteLogConfiguration request param
type DeleteLogConfigurationParam struct {
	BaseParam
	Params DeleteLogConfigurationDetailParam `json:"params"`
}
