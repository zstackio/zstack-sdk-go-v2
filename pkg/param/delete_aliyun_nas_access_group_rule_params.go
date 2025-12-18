// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunNasAccessGroupRuleDetailParam DeleteAliyunNasAccessGroupRule detail param
type DeleteAliyunNasAccessGroupRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunNasAccessGroupRuleParam DeleteAliyunNasAccessGroupRule request param
type DeleteAliyunNasAccessGroupRuleParam struct {
	BaseParam
	Params DeleteAliyunNasAccessGroupRuleDetailParam `json:"params"`
}
