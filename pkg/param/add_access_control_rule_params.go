// Copyright (c) ZStack.io, Inc.

package param

// AddAccessControlRuleDetailParam AddAccessControlRule详细参数
type AddAccessControlRuleDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"rule" validate:"required"` // 必填
	rest string `json:"controlStrategy" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddAccessControlRuleParam AddAccessControlRule请求参数
type AddAccessControlRuleParam struct {
	BaseParam
	Params AddAccessControlRuleDetailParam `json:"params"` // 详细参数
}

