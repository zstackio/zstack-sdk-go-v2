// Copyright (c) ZStack.io, Inc.

package param

// GetZBoxBackupDetailsDetailParam GetZBoxBackupDetails detail param
type GetZBoxBackupDetailsDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetZBoxBackupDetailsParam GetZBoxBackupDetails request param
type GetZBoxBackupDetailsParam struct {
	BaseParam
	Params GetZBoxBackupDetailsDetailParam `json:"params"`
}
