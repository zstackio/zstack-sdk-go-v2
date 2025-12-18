// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SchedulerJobGroupJobRefInventoryView SchedulerJobGroupJobRef
type SchedulerJobGroupJobRefInventoryView struct {
	rest string `json:"schedulerJobGroupUuid,omitempty"`
	rest string `json:"schedulerJobUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

