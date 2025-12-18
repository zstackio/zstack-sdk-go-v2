// Copyright (c) ZStack.io, Inc.

package param

// GetBareMetal2ChassisPowerStatusDetailParam GetBareMetal2ChassisPowerStatus详细参数
type GetBareMetal2ChassisPowerStatusDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetBareMetal2ChassisPowerStatusParam GetBareMetal2ChassisPowerStatus请求参数
type GetBareMetal2ChassisPowerStatusParam struct {
	BaseParam
	Params GetBareMetal2ChassisPowerStatusDetailParam `json:"params"` // 详细参数
}

