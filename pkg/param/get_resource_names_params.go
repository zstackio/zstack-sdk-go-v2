// Copyright (c) ZStack.io, Inc.

package param

// GetResourceNamesDetailParam GetResourceNames detail param
type GetResourceNamesDetailParam struct {
	Uuids []string `json:"uuids" validate:"required"`
}

// GetResourceNamesParam GetResourceNames request param
type GetResourceNamesParam struct {
	BaseParam
	Params GetResourceNamesDetailParam `json:"params"`
}
