// Copyright (c) ZStack.io, Inc.

package param

// UpdateSchedulerTriggerDetailParam UpdateSchedulerTrigger detail param
type UpdateSchedulerTriggerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	SchedulerInterval int `json:"schedulerInterval,omitempty"`
	RepeatCount int `json:"repeatCount,omitempty"`
	StartTime int64 `json:"startTime,omitempty"`
	Cron string `json:"cron,omitempty"`
	SchedulerType string `json:"schedulerType,omitempty"`
}

// UpdateSchedulerTriggerParam UpdateSchedulerTrigger request param
type UpdateSchedulerTriggerParam struct {
	BaseParam
	Params UpdateSchedulerTriggerDetailParam `json:"params"`
}
