// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MonitorGroupAlarmInventoryView MonitorGroupAlarm
type MonitorGroupAlarmInventoryView struct {
	rest string `json:"groupUuid,omitempty"`
	rest string `json:"alarmUuid,omitempty"`
	rest string `json:"metricRuleTemplateUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest string `json:"uuid,omitempty"`
}

