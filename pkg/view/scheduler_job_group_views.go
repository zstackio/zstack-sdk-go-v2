// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SchedulerJobGroupInventoryView SchedulerJobGroup
type SchedulerJobGroupInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	JobType string `json:"jobType,omitempty"`
	JobData string `json:"jobData,omitempty"`
	TriggersUuid []string `json:"triggersUuid,omitempty"`
	JobsUuid []string `json:"jobsUuid,omitempty"`
}

