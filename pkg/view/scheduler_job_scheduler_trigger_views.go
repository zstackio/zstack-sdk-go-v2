// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SchedulerJobSchedulerTriggerInventoryView SchedulerJobSchedulerTrigger
type SchedulerJobSchedulerTriggerInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	SchedulerJobUuid string `json:"schedulerJobUuid,omitempty"`
	SchedulerTriggerUuid string `json:"schedulerTriggerUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

