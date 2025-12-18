// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SchedulerJobHistoryInventoryView SchedulerJobHistory
type SchedulerJobHistoryInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"triggerUuid,omitempty"`
	rest string `json:"schedulerJobUuid,omitempty"`
	rest string `json:"schedulerJobGroupUuid,omitempty"`
	rest string `json:"jobType,omitempty"`
	rest time.Time `json:"startTime,omitempty"`
	rest int64 `json:"executeTime,omitempty"`
	rest string `json:"targetResourceUuid,omitempty"`
	rest string `json:"requestDump,omitempty"`
	rest string `json:"resultDump,omitempty"`
	rest bool `json:"success,omitempty"`
	rest string `json:"fireInstanceId,omitempty"`
}

