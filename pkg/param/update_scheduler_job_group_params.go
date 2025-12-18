// Copyright (c) ZStack.io, Inc.

package param

// UpdateSchedulerJobGroupDetailParam UpdateSchedulerJobGroup detail param
type UpdateSchedulerJobGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

// UpdateSchedulerJobGroupParam UpdateSchedulerJobGroup request param
type UpdateSchedulerJobGroupParam struct {
	BaseParam
	Params UpdateSchedulerJobGroupDetailParam `json:"params"`
}
