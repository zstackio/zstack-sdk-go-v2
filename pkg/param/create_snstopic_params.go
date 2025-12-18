// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSTopicDetailParam CreateSNSTopic detail param
type CreateSNSTopicDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Locale string `json:"locale,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSTopicParam CreateSNSTopic request param
type CreateSNSTopicParam struct {
	BaseParam
	Params CreateSNSTopicDetailParam `json:"params"`
}
