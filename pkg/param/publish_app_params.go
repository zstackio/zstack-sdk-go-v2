// Copyright (c) ZStack.io, Inc.

package param

// PublishAppDetailParam PublishApp详细参数
type PublishAppDetailParam struct {
	rest string `json:"buildAppUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"parameters,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// PublishAppParam PublishApp请求参数
type PublishAppParam struct {
	BaseParam
	Params PublishAppDetailParam `json:"params"` // 详细参数
}

