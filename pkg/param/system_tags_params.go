// Copyright (c) ZStack.io, Inc.

package param

// CreateSystemTagsDetailParam CreateSystemTags详细参数
type CreateSystemTagsDetailParam struct {
	rest string `json:"resourceType" validate:"required"` // 必填
	rest string `json:"resourceUuid" validate:"required"` // 必填
	rest []string `json:"tags" validate:"required"` // 必填
}

// CreateSystemTagsParam CreateSystemTags请求参数
type CreateSystemTagsParam struct {
	BaseParam
	Params CreateSystemTagsDetailParam `json:"params"` // 详细参数
}

