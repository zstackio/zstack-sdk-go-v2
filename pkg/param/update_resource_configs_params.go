// Copyright (c) ZStack.io, Inc.

package param

// UpdateResourceConfigsDetailParam UpdateResourceConfigs detail param
type UpdateResourceConfigsDetailParam struct {
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	ResourceConfigs []ResourceConfigAOParam `json:"resourceConfigs" validate:"required"`
}

// UpdateResourceConfigsParam UpdateResourceConfigs request param
type UpdateResourceConfigsParam struct {
	BaseParam
	Params UpdateResourceConfigsDetailParam `json:"params"`
}
