// Copyright (c) ZStack.io, Inc.

package param

// GetBaremetalChassisPowerStatusDetailParam GetBaremetalChassisPowerStatus detail param
type GetBaremetalChassisPowerStatusDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetBaremetalChassisPowerStatusParam GetBaremetalChassisPowerStatus request param
type GetBaremetalChassisPowerStatusParam struct {
	BaseParam
	Params GetBaremetalChassisPowerStatusDetailParam `json:"params"`
}
