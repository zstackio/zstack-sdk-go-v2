// Copyright (c) ZStack.io, Inc.

package param

// ChangeVmNicStateDetailParam ChangeVmNicState详细参数
type ChangeVmNicStateDetailParam struct {
	rest string `json:"vmNicUuid" validate:"required"` // 必填
	rest string `json:"state" validate:"required"` // 必填
}

// ChangeVmNicStateParam ChangeVmNicState请求参数
type ChangeVmNicStateParam struct {
	BaseParam
	Params ChangeVmNicStateDetailParam `json:"params"` // 详细参数
}

