// Copyright (c) ZStack.io, Inc.

package param

// ChangeResourceOwnerDetailParam ChangeResourceOwner detail param
type ChangeResourceOwnerDetailParam struct {
	AccountUuid string `json:"accountUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid" validate:"required"`
}

// ChangeResourceOwnerParam ChangeResourceOwner request param
type ChangeResourceOwnerParam struct {
	BaseParam
	Params ChangeResourceOwnerDetailParam `json:"params"`
}
