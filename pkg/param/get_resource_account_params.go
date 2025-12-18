// Copyright (c) ZStack.io, Inc.

package param

// GetResourceAccountDetailParam GetResourceAccount detail param
type GetResourceAccountDetailParam struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
}

// GetResourceAccountParam GetResourceAccount request param
type GetResourceAccountParam struct {
	BaseParam
	Params GetResourceAccountDetailParam `json:"params"`
}
