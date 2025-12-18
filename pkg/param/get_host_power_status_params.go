// Copyright (c) ZStack.io, Inc.

package param

// GetHostPowerStatusDetailParam GetHostPowerStatus detail param
type GetHostPowerStatusDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Method string `json:"method,omitempty"`
}

// GetHostPowerStatusParam GetHostPowerStatus request param
type GetHostPowerStatusParam struct {
	BaseParam
	Params GetHostPowerStatusDetailParam `json:"params"`
}
