// Copyright (c) ZStack.io, Inc.

package param

// ChangeResourceOwnerDetailParam ChangeResourceOwner详细参数
type ChangeResourceOwnerDetailParam struct {
	rest string `json:"accountUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid" validate:"required"` // 必填
}

// ChangeResourceOwnerParam ChangeResourceOwner请求参数
type ChangeResourceOwnerParam struct {
	BaseParam
	Params ChangeResourceOwnerDetailParam `json:"params"` // 详细参数
}

