// Copyright (c) ZStack.io, Inc.

package param

// ChangeVfNicHaStateDetailParam ChangeVfNicHaState detail param
type ChangeVfNicHaStateDetailParam struct {
	VfNicUuid string `json:"vfNicUuid" validate:"required"`
	HaState string `json:"haState" validate:"required"`
}

// ChangeVfNicHaStateParam ChangeVfNicHaState request param
type ChangeVfNicHaStateParam struct {
	BaseParam
	Params ChangeVfNicHaStateDetailParam `json:"params"`
}
