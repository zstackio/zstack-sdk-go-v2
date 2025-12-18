// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SchedulerJobGroupSchedulerTriggerRefInventoryView SchedulerJobGroupSchedulerTriggerRef
type SchedulerJobGroupSchedulerTriggerRefInventoryView struct {
	SchedulerJobGroupUuid string `json:"schedulerJobGroupUuid,omitempty"`
	SchedulerTriggerUuid string `json:"schedulerTriggerUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

