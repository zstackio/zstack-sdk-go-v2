// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SchedulerJobSchedulerTriggerInventoryView SchedulerJobSchedulerTrigger
type SchedulerJobSchedulerTriggerInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"schedulerJobUuid,omitempty"`
	rest string `json:"schedulerTriggerUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

