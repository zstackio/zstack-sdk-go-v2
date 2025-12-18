// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AlarmInventoryView Alarm
type AlarmInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"comparisonOperator,omitempty"`
	rest int `json:"period,omitempty"`
	rest string `json:"namespace,omitempty"`
	rest string `json:"metricName,omitempty"`
	rest float64 `json:"threshold,omitempty"`
	rest int `json:"repeatInterval,omitempty"`
	rest int `json:"repeatCount,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"state,omitempty"`
	rest bool `json:"enableRecovery,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []AlarmLabelInventoryView `json:"labels,omitempty"`
	rest []AlarmActionInventoryView `json:"actions,omitempty"`
	rest string `json:"emergencyLevel,omitempty"`
}

