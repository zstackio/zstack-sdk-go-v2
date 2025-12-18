// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SchedulerTriggerInventoryView SchedulerTrigger
type SchedulerTriggerInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"cron,omitempty"`
	rest string `json:"schedulerType,omitempty"`
	rest int `json:"schedulerInterval,omitempty"`
	rest int `json:"repeatCount,omitempty"`
	rest time.Time `json:"startTime,omitempty"`
	rest time.Time `json:"stopTime,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []string `json:"jobsUuid,omitempty"`
	rest []string `json:"jobGroupsUuid,omitempty"`
}

