// Copyright (c) ZStack.io, Inc.

package param

// DeleteSchedulerTriggerDetailParam DeleteSchedulerTrigger detail param
type DeleteSchedulerTriggerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSchedulerTriggerParam DeleteSchedulerTrigger request param
type DeleteSchedulerTriggerParam struct {
	BaseParam
	Params DeleteSchedulerTriggerDetailParam `json:"params"`
}
