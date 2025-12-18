// Copyright (c) ZStack.io, Inc.

package param

// CreateResourceStackFromAppDetailParam CreateResourceStackFromApp detail param
type CreateResourceStackFromAppDetailParam struct {
	AppUuid string `json:"appUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Rollback bool `json:"rollback,omitempty"`
	Parameters string `json:"parameters,omitempty"`
	WithoutAppInfo bool `json:"withoutAppInfo,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateResourceStackFromAppParam CreateResourceStackFromApp request param
type CreateResourceStackFromAppParam struct {
	BaseParam
	Params CreateResourceStackFromAppDetailParam `json:"params"`
}
