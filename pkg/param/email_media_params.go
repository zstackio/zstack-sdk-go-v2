// Copyright (c) ZStack.io, Inc.

package param

// CreateEmailMediaDetailParam CreateEmailMedia详细参数
type CreateEmailMediaDetailParam struct {
	rest string `json:"smtpServer" validate:"required"` // 必填
	rest int `json:"smtpPort" validate:"required"` // 必填
	rest string `json:"username,omitempty"`
	rest string `json:"password,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateEmailMediaParam CreateEmailMedia请求参数
type CreateEmailMediaParam struct {
	BaseParam
	Params CreateEmailMediaDetailParam `json:"params"` // 详细参数
}

