// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SchedulerJobGroupJobRefInventoryView SchedulerJobGroupJobRef
type SchedulerJobGroupJobRefInventoryView struct {
	SchedulerJobGroupUuid *string `json:"schedulerJobGroupUuid,omitempty"`
	SchedulerJobUuid *string `json:"schedulerJobUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// AddSchedulerJobsToSchedulerJobGroupEventView AddSchedulerJobsToSchedulerJobGroupEvent
type AddSchedulerJobsToSchedulerJobGroupEventView struct {
	Inventories []SchedulerJobGroupJobRefInventoryView `json:"inventories,omitempty"`
}

