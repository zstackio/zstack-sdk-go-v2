// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateUserTagParamDetail CreateUserTag detail param
type CreateUserTagParamDetail struct {
	ResourceType string `json:"resourceType" validate:"required"`
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	Tag string `json:"tag" validate:"required"`
}

// CreateUserTagParam CreateUserTag request param
type CreateUserTagParam struct {
	BaseParam
	Params CreateUserTagParamDetail `json:"params"`
}
