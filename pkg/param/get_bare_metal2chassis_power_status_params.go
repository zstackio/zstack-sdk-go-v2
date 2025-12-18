// Copyright (c) ZStack.io, Inc.

package param

// GetBareMetal2ChassisPowerStatusDetailParam GetBareMetal2ChassisPowerStatus detail param
type GetBareMetal2ChassisPowerStatusDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetBareMetal2ChassisPowerStatusParam GetBareMetal2ChassisPowerStatus request param
type GetBareMetal2ChassisPowerStatusParam struct {
	BaseParam
	Params GetBareMetal2ChassisPowerStatusDetailParam `json:"params"`
}
