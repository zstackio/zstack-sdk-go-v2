// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SchedulerTriggerInventoryView SchedulerTrigger
type SchedulerTriggerInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Cron string `json:"cron,omitempty"`
	SchedulerType string `json:"schedulerType,omitempty"`
	SchedulerInterval int `json:"schedulerInterval,omitempty"`
	RepeatCount int `json:"repeatCount,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	StopTime time.Time `json:"stopTime,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	JobsUuid []string `json:"jobsUuid,omitempty"`
	JobGroupsUuid []string `json:"jobGroupsUuid,omitempty"`
}

