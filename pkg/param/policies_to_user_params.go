// Copyright (c) ZStack.io, Inc.

package param

// AttachPoliciesToUserDetailParam AttachPoliciesToUser详细参数
type AttachPoliciesToUserDetailParam struct {
	rest string `json:"userUuid" validate:"required"` // 必填
	rest []string `json:"policyUuids" validate:"required"` // 必填
}

// AttachPoliciesToUserParam AttachPoliciesToUser请求参数
type AttachPoliciesToUserParam struct {
	BaseParam
	Params AttachPoliciesToUserDetailParam `json:"params"` // 详细参数
}

