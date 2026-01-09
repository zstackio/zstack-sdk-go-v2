// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SchedulerJobInventoryView SchedulerJob
type SchedulerJobInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	TargetResourceUuid *string `json:"targetResourceUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	State *string `json:"state,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	JobData *string `json:"jobData,omitempty"`
	JobClassName *string `json:"jobClassName,omitempty"`
	TriggersUuid []string `json:"triggersUuid,omitempty"`
	SchedulerJobGroupUuids []string `json:"schedulerJobGroupUuids,omitempty"`
}

// GetNoTriggerSchedulerJobsView GetNoTriggerSchedulerJobs
type GetNoTriggerSchedulerJobsView struct {
	Inventories []SchedulerJobInventoryView `json:"inventories,omitempty"`
}

// CreateSchedulerJobEventView CreateSchedulerJobEvent
type CreateSchedulerJobEventView struct {
	Inventory SchedulerJobInventoryView `json:"inventory,omitempty"`
}

// UpdateSchedulerJobEventView UpdateSchedulerJobEvent
type UpdateSchedulerJobEventView struct {
	Inventory SchedulerJobInventoryView `json:"inventory,omitempty"`
}

// ChangeSchedulerStateEventView ChangeSchedulerStateEvent
type ChangeSchedulerStateEventView struct {
	Inventory SchedulerJobInventoryView `json:"inventory,omitempty"`
}

// DeleteSchedulerJobEventView DeleteSchedulerJobEvent
type DeleteSchedulerJobEventView struct {
	Success bool `json:"success,omitempty"`
}

// QuerySchedulerJobView QuerySchedulerJob
type QuerySchedulerJobView struct {
	Inventories []SchedulerJobInventoryView `json:"inventories,omitempty"`
}

