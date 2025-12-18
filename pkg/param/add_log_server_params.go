// Copyright (c) ZStack.io, Inc.

package param

// AddLogServerDetailParam AddLogServer详细参数
type AddLogServerDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"category" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"level,omitempty"`
	rest string `json:"configuration" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddLogServerParam AddLogServer请求参数
type AddLogServerParam struct {
	BaseParam
	Params AddLogServerDetailParam `json:"params"` // 详细参数
}

