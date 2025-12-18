// Copyright (c) ZStack.io, Inc.

package param

// DeleteNvmeServerDetailParam DeleteNvmeServer detail param
type DeleteNvmeServerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteNvmeServerParam DeleteNvmeServer request param
type DeleteNvmeServerParam struct {
	BaseParam
	Params DeleteNvmeServerDetailParam `json:"params"`
}
