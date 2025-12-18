// Copyright (c) ZStack.io, Inc.

package param

// ChangeBackupStorageStateDetailParam ChangeBackupStorageState detail param
type ChangeBackupStorageStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeBackupStorageStateParam ChangeBackupStorageState request param
type ChangeBackupStorageStateParam struct {
	BaseParam
	Params ChangeBackupStorageStateDetailParam `json:"params"`
}
