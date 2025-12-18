// Copyright (c) ZStack.io, Inc.

package param

// AddAccessControlRuleDetailParam AddAccessControlRule detail param
type AddAccessControlRuleDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Rule string `json:"rule" validate:"required"`
	ControlStrategy string `json:"controlStrategy" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAccessControlRuleParam AddAccessControlRule request param
type AddAccessControlRuleParam struct {
	BaseParam
	Params AddAccessControlRuleDetailParam `json:"params"`
}
