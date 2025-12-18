// Copyright (c) ZStack.io, Inc.

package param

// ChangeBackupStorageStateDetailParam ChangeBackupStorageState详细参数
type ChangeBackupStorageStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeBackupStorageStateParam ChangeBackupStorageState请求参数
type ChangeBackupStorageStateParam struct {
	BaseParam
	Params ChangeBackupStorageStateDetailParam `json:"params"` // 详细参数
}

