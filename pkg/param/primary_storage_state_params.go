// Copyright (c) ZStack.io, Inc.

package param

// ChangePrimaryStorageStateDetailParam ChangePrimaryStorageState详细参数
type ChangePrimaryStorageStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangePrimaryStorageStateParam ChangePrimaryStorageState请求参数
type ChangePrimaryStorageStateParam struct {
	BaseParam
	Params ChangePrimaryStorageStateDetailParam `json:"params"` // 详细参数
}

