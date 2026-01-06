// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AlarmActionInventoryView AlarmAction
type AlarmActionInventoryView struct {
	AlarmUuid string `json:"alarmUuid,omitempty"`
	ActionType string `json:"actionType,omitempty"`
	ActionUuid string `json:"actionUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

