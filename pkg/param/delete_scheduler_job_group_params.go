// Copyright (c) ZStack.io, Inc.

package param

// DeleteSchedulerJobGroupDetailParam DeleteSchedulerJobGroup detail param
type DeleteSchedulerJobGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSchedulerJobGroupParam DeleteSchedulerJobGroup request param
type DeleteSchedulerJobGroupParam struct {
	BaseParam
	Params DeleteSchedulerJobGroupDetailParam `json:"params"`
}
