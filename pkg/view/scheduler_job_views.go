// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SchedulerJobInventoryView SchedulerJob
type SchedulerJobInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"targetResourceUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"jobData,omitempty"`
	rest string `json:"jobClassName,omitempty"`
	rest []string `json:"triggersUuid,omitempty"`
	rest []string `json:"schedulerJobGroupUuids,omitempty"`
}

