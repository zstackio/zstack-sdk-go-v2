// Copyright (c) ZStack.io, Inc.

package param

// CreateSchedulerTriggerDetailParam CreateSchedulerTrigger detail param
type CreateSchedulerTriggerDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	SchedulerInterval int `json:"schedulerInterval,omitempty"`
	RepeatCount int `json:"repeatCount,omitempty"`
	StartTime int64 `json:"startTime,omitempty"`
	SchedulerType string `json:"schedulerType" validate:"required"`
	Cron string `json:"cron,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSchedulerTriggerParam CreateSchedulerTrigger request param
type CreateSchedulerTriggerParam struct {
	BaseParam
	Params CreateSchedulerTriggerDetailParam `json:"params"`
}
