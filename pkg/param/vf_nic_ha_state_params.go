// Copyright (c) ZStack.io, Inc.

package param

// ChangeVfNicHaStateDetailParam ChangeVfNicHaState详细参数
type ChangeVfNicHaStateDetailParam struct {
	rest string `json:"vfNicUuid" validate:"required"` // 必填
	rest string `json:"haState" validate:"required"` // 必填
}

// ChangeVfNicHaStateParam ChangeVfNicHaState请求参数
type ChangeVfNicHaStateParam struct {
	BaseParam
	Params ChangeVfNicHaStateDetailParam `json:"params"` // 详细参数
}

