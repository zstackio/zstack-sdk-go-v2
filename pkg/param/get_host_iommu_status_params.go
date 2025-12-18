// Copyright (c) ZStack.io, Inc.

package param

// GetHostIommuStatusDetailParam GetHostIommuStatus detail param
type GetHostIommuStatusDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetHostIommuStatusParam GetHostIommuStatus request param
type GetHostIommuStatusParam struct {
	BaseParam
	Params GetHostIommuStatusDetailParam `json:"params"`
}
