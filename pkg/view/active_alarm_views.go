// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ActiveAlarmInventoryView ActiveAlarm
type ActiveAlarmInventoryView struct {
	rest string `json:"templateUuid,omitempty"`
	rest string `json:"alarmUuid,omitempty"`
	rest string `json:"namespace,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest string `json:"uuid,omitempty"`
}

