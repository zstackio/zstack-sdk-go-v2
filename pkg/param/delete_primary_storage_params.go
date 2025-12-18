// Copyright (c) ZStack.io, Inc.

package param

// DeletePrimaryStorageDetailParam DeletePrimaryStorage detail param
type DeletePrimaryStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePrimaryStorageParam DeletePrimaryStorage request param
type DeletePrimaryStorageParam struct {
	BaseParam
	Params DeletePrimaryStorageDetailParam `json:"params"`
}
