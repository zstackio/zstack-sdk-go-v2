// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SchedulerJobSchedulerTriggerInventoryView SchedulerJobSchedulerTrigger
type SchedulerJobSchedulerTriggerInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	SchedulerJobUuid string `json:"schedulerJobUuid,omitempty"`
	SchedulerTriggerUuid string `json:"schedulerTriggerUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// AddSchedulerJobToSchedulerTriggerEventView AddSchedulerJobToSchedulerTriggerEvent
type AddSchedulerJobToSchedulerTriggerEventView struct {
	Inventory SchedulerJobSchedulerTriggerInventoryView `json:"inventory,omitempty"`
}

