// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AlarmDataAckInventoryView AlarmDataAck
type AlarmDataAckInventoryView struct {
	rest string `json:"alarmUuid,omitempty"`
	rest string `json:"alertDataUuid,omitempty"`
	rest string `json:"alertType,omitempty"`
	rest int64 `json:"ackPeriod,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest time.Time `json:"ackDate,omitempty"`
	rest bool `json:"resumeAlert,omitempty"`
	rest string `json:"operatorAccountUuid,omitempty"`
}

