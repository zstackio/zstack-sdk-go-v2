// Copyright (c) ZStack.io, Inc.

package param

// ApplyRuleSetChangesDetailParam ApplyRuleSetChanges详细参数
type ApplyRuleSetChangesDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ApplyRuleSetChangesParam ApplyRuleSetChanges请求参数
type ApplyRuleSetChangesParam struct {
	BaseParam
	Params ApplyRuleSetChangesDetailParam `json:"params"` // 详细参数
}

