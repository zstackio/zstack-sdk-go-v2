// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SchedulerJobGroupInventoryView SchedulerJobGroup
type SchedulerJobGroupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"jobType,omitempty"`
	rest string `json:"jobData,omitempty"`
	rest []string `json:"triggersUuid,omitempty"`
	rest []string `json:"jobsUuid,omitempty"`
}

