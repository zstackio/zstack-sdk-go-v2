// Copyright (c) ZStack.io, Inc.

package param

// ShareResourceDetailParam ShareResource详细参数
type ShareResourceDetailParam struct {
	rest []string `json:"resourceUuids" validate:"required"` // 必填
	rest []string `json:"accountUuids,omitempty"`
	rest bool `json:"toPublic,omitempty"`
}

// ShareResourceParam ShareResource请求参数
type ShareResourceParam struct {
	BaseParam
	Params ShareResourceDetailParam `json:"params"` // 详细参数
}

