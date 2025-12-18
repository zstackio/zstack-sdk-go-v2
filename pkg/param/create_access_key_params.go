// Copyright (c) ZStack.io, Inc.

package param

// CreateAccessKeyDetailParam CreateAccessKey detail param
type CreateAccessKeyDetailParam struct {
	AccountUuid string `json:"accountUuid" validate:"required"`
	UserUuid string `json:"userUuid" validate:"required"`
	Description string `json:"description,omitempty"`
	AccessKeyID string `json:"AccessKeyID,omitempty"`
	AccessKeySecret string `json:"AccessKeySecret,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAccessKeyParam CreateAccessKey request param
type CreateAccessKeyParam struct {
	BaseParam
	Params CreateAccessKeyDetailParam `json:"params"`
}
