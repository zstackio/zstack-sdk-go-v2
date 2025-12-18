// Copyright (c) ZStack.io, Inc.

package param

// BindModelToServiceDetailParam BindModelToService detail param
type BindModelToServiceDetailParam struct {
	ModelUuid string `json:"modelUuid" validate:"required"`
	ModelServiceUuid string `json:"modelServiceUuid" validate:"required"`
}

// BindModelToServiceParam BindModelToService request param
type BindModelToServiceParam struct {
	BaseParam
	Params BindModelToServiceDetailParam `json:"params"`
}
