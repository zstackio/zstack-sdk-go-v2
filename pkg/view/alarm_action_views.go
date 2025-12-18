// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AlarmActionInventoryView AlarmAction
type AlarmActionInventoryView struct {
	rest string `json:"alarmUuid,omitempty"`
	rest string `json:"actionType,omitempty"`
	rest string `json:"actionUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

