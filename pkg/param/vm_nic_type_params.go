// Copyright (c) ZStack.io, Inc.

package param

// ChangeVmNicTypeDetailParam ChangeVmNicType详细参数
type ChangeVmNicTypeDetailParam struct {
	rest string `json:"vmNicUuid" validate:"required"` // 必填
	rest string `json:"vmNicType" validate:"required"` // 必填
}

// ChangeVmNicTypeParam ChangeVmNicType请求参数
type ChangeVmNicTypeParam struct {
	BaseParam
	Params ChangeVmNicTypeDetailParam `json:"params"` // 详细参数
}

