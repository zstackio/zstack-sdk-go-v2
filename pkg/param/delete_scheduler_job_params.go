// Copyright (c) ZStack.io, Inc.

package param

// DeleteSchedulerJobDetailParam DeleteSchedulerJob detail param
type DeleteSchedulerJobDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSchedulerJobParam DeleteSchedulerJob request param
type DeleteSchedulerJobParam struct {
	BaseParam
	Params DeleteSchedulerJobDetailParam `json:"params"`
}
