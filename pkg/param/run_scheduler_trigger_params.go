// Copyright (c) ZStack.io, Inc.

package param

// RunSchedulerTriggerDetailParam RunSchedulerTrigger detail param
type RunSchedulerTriggerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	JobUuids []string `json:"jobUuids,omitempty"`
}

// RunSchedulerTriggerParam RunSchedulerTrigger request param
type RunSchedulerTriggerParam struct {
	BaseParam
	Params RunSchedulerTriggerDetailParam `json:"params"`
}
