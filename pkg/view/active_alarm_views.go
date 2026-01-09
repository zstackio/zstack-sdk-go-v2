// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ActiveAlarmInventoryView ActiveAlarm
type ActiveAlarmInventoryView struct {
	TemplateUuid *string `json:"templateUuid,omitempty"`
	AlarmUuid *string `json:"alarmUuid,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

// QueryActiveAlarmView QueryActiveAlarm
type QueryActiveAlarmView struct {
	Inventories []ActiveAlarmInventoryView `json:"inventories,omitempty"`
}

