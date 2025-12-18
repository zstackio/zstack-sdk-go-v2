// Copyright (c) ZStack.io, Inc.

package param

// UnbindModelFromServiceDetailParam UnbindModelFromService detail param
type UnbindModelFromServiceDetailParam struct {
	ModelUuid string `json:"modelUuid" validate:"required"`
	ModelServiceUuid string `json:"modelServiceUuid" validate:"required"`
}

// UnbindModelFromServiceParam UnbindModelFromService request param
type UnbindModelFromServiceParam struct {
	BaseParam
	Params UnbindModelFromServiceDetailParam `json:"params"`
}
