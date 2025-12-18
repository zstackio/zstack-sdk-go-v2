// Copyright (c) ZStack.io, Inc.

package param

// GetHostPowerStatusDetailParam GetHostPowerStatus详细参数
type GetHostPowerStatusDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"method,omitempty"`
}

// GetHostPowerStatusParam GetHostPowerStatus请求参数
type GetHostPowerStatusParam struct {
	BaseParam
	Params GetHostPowerStatusDetailParam `json:"params"` // 详细参数
}

