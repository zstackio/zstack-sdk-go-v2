// Copyright (c) ZStack.io, Inc.

package param

// ChangeVmNicStateDetailParam ChangeVmNicState detail param
type ChangeVmNicStateDetailParam struct {
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
	State string `json:"state" validate:"required"`
}

// ChangeVmNicStateParam ChangeVmNicState request param
type ChangeVmNicStateParam struct {
	BaseParam
	Params ChangeVmNicStateDetailParam `json:"params"`
}
