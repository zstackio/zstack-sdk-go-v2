// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SchedulerJobGroupSchedulerTriggerRefInventoryView SchedulerJobGroupSchedulerTriggerRef
type SchedulerJobGroupSchedulerTriggerRefInventoryView struct {
	rest string `json:"schedulerJobGroupUuid,omitempty"`
	rest string `json:"schedulerTriggerUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

