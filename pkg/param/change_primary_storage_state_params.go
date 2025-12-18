// Copyright (c) ZStack.io, Inc.

package param

// ChangePrimaryStorageStateDetailParam ChangePrimaryStorageState detail param
type ChangePrimaryStorageStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangePrimaryStorageStateParam ChangePrimaryStorageState request param
type ChangePrimaryStorageStateParam struct {
	BaseParam
	Params ChangePrimaryStorageStateDetailParam `json:"params"`
}
