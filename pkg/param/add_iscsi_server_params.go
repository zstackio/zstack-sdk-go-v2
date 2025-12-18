// Copyright (c) ZStack.io, Inc.

package param

// AddIscsiServerDetailParam AddIscsiServer详细参数
type AddIscsiServerDetailParam struct {
	rest string `json:"name,omitempty"`
	rest string `json:"ip" validate:"required"` // 必填
	rest int `json:"port,omitempty"`
	rest string `json:"chapUserName,omitempty"`
	rest string `json:"chapUserPassword,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddIscsiServerParam AddIscsiServer请求参数
type AddIscsiServerParam struct {
	BaseParam
	Params AddIscsiServerDetailParam `json:"params"` // 详细参数
}

