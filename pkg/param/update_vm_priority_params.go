// Copyright (c) ZStack.io, Inc.

package param

// UpdateVmPriorityDetailParam UpdateVmPriority detail param
type UpdateVmPriorityDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Priority string `json:"priority" validate:"required"`
}

// UpdateVmPriorityParam UpdateVmPriority request param
type UpdateVmPriorityParam struct {
	BaseParam
	Params UpdateVmPriorityDetailParam `json:"params"`
}
