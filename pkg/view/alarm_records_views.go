// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AlarmRecordsInventoryView AlarmRecords
type AlarmRecordsInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest int64 `json:"createTime,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"alarmName,omitempty"`
	rest string `json:"alarmStatus,omitempty"`
	rest string `json:"alarmUuid,omitempty"`
	rest string `json:"comparisonOperator,omitempty"`
	rest string `json:"context,omitempty"`
	rest string `json:"dataUuid,omitempty"`
	rest string `json:"emergencyLevel,omitempty"`
	rest string `json:"labels,omitempty"`
	rest string `json:"metricName,omitempty"`
	rest float64 `json:"metricValue,omitempty"`
	rest string `json:"namespace,omitempty"`
	rest int `json:"period,omitempty"`
	rest bool `json:"readStatus,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest float64 `json:"threshold,omitempty"`
}

