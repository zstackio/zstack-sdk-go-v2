// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MonitorGroupAlarmInventoryView MonitorGroupAlarm
type MonitorGroupAlarmInventoryView struct {
	GroupUuid string `json:"groupUuid,omitempty"`
	AlarmUuid string `json:"alarmUuid,omitempty"`
	MetricRuleTemplateUuid string `json:"metricRuleTemplateUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

// QueryMonitorGroupAlarmView QueryMonitorGroupAlarm
type QueryMonitorGroupAlarmView struct {
	Inventories []MonitorGroupAlarmInventoryView `json:"inventories,omitempty"`
}

