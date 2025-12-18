// Copyright (c) ZStack.io, Inc.

package param

// DeleteModelServiceDetailParam DeleteModelService detail param
type DeleteModelServiceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteModelServiceParam DeleteModelService request param
type DeleteModelServiceParam struct {
	BaseParam
	Params DeleteModelServiceDetailParam `json:"params"`
}
