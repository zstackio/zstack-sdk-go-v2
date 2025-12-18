// Copyright (c) ZStack.io, Inc.

package param

// GetResourceStackVmStatusDetailParam GetResourceStackVmStatus detail param
type GetResourceStackVmStatusDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetResourceStackVmStatusParam GetResourceStackVmStatus request param
type GetResourceStackVmStatusParam struct {
	BaseParam
	Params GetResourceStackVmStatusDetailParam `json:"params"`
}
