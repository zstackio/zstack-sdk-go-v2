// Copyright (c) ZStack.io, Inc.

package param

// GetBaremetalChassisPowerStatusDetailParam GetBaremetalChassisPowerStatus详细参数
type GetBaremetalChassisPowerStatusDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetBaremetalChassisPowerStatusParam GetBaremetalChassisPowerStatus请求参数
type GetBaremetalChassisPowerStatusParam struct {
	BaseParam
	Params GetBaremetalChassisPowerStatusDetailParam `json:"params"` // 详细参数
}

