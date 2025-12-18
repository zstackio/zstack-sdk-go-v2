// Copyright (c) ZStack.io, Inc.

package param

// ApplyRuleSetChangesDetailParam ApplyRuleSetChanges detail param
type ApplyRuleSetChangesDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ApplyRuleSetChangesParam ApplyRuleSetChanges request param
type ApplyRuleSetChangesParam struct {
	BaseParam
	Params ApplyRuleSetChangesDetailParam `json:"params"`
}
