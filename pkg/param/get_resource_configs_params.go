// Copyright (c) ZStack.io, Inc.

package param

// GetResourceConfigsDetailParam GetResourceConfigs detail param
type GetResourceConfigsDetailParam struct {
	Category string `json:"category" validate:"required"`
	Names []string `json:"names" validate:"required"`
	ResourceUuid string `json:"resourceUuid" validate:"required"`
}

// GetResourceConfigsParam GetResourceConfigs request param
type GetResourceConfigsParam struct {
	BaseParam
	Params GetResourceConfigsDetailParam `json:"params"`
}
