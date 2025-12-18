// Copyright (c) ZStack.io, Inc.

package param

// UpdateSchedulerJobDetailParam UpdateSchedulerJob detail param
type UpdateSchedulerJobDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

// UpdateSchedulerJobParam UpdateSchedulerJob request param
type UpdateSchedulerJobParam struct {
	BaseParam
	Params UpdateSchedulerJobDetailParam `json:"params"`
}
