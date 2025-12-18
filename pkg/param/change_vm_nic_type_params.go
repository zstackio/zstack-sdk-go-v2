// Copyright (c) ZStack.io, Inc.

package param

// ChangeVmNicTypeDetailParam ChangeVmNicType detail param
type ChangeVmNicTypeDetailParam struct {
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
	VmNicType string `json:"vmNicType" validate:"required"`
}

// ChangeVmNicTypeParam ChangeVmNicType request param
type ChangeVmNicTypeParam struct {
	BaseParam
	Params ChangeVmNicTypeDetailParam `json:"params"`
}
