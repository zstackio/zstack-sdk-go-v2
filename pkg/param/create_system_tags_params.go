// Copyright (c) ZStack.io, Inc.

package param

// CreateSystemTagsDetailParam CreateSystemTags detail param
type CreateSystemTagsDetailParam struct {
	ResourceType string `json:"resourceType" validate:"required"`
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	Tags []string `json:"tags" validate:"required"`
}

// CreateSystemTagsParam CreateSystemTags request param
type CreateSystemTagsParam struct {
	BaseParam
	Params CreateSystemTagsDetailParam `json:"params"`
}
