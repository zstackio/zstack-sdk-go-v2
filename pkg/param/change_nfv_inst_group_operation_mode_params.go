// Copyright (c) ZStack.io, Inc.

package param

// ChangeNfvInstGroupOperationModeDetailParam ChangeNfvInstGroupOperationMode detail param
type ChangeNfvInstGroupOperationModeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	OperationMode string `json:"operationMode" validate:"required"`
}

// ChangeNfvInstGroupOperationModeParam ChangeNfvInstGroupOperationMode request param
type ChangeNfvInstGroupOperationModeParam struct {
	BaseParam
	Params ChangeNfvInstGroupOperationModeDetailParam `json:"params"`
}
