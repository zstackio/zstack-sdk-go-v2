// Copyright (c) ZStack.io, Inc.

package param

// ReconnectPrimaryStorageDetailParam ReconnectPrimaryStorage detail param
type ReconnectPrimaryStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectPrimaryStorageParam ReconnectPrimaryStorage request param
type ReconnectPrimaryStorageParam struct {
	BaseParam
	Params ReconnectPrimaryStorageDetailParam `json:"params"`
}
