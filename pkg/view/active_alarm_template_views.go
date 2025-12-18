// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ActiveAlarmTemplateInventoryView ActiveAlarmTemplate
type ActiveAlarmTemplateInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"alarmName,omitempty"`
	rest string `json:"comparisonOperator,omitempty"`
	rest int `json:"period,omitempty"`
	rest int `json:"repeatInterval,omitempty"`
	rest int `json:"repeatCount,omitempty"`
	rest string `json:"namespace,omitempty"`
	rest string `json:"metricName,omitempty"`
	rest float64 `json:"threshold,omitempty"`
	rest string `json:"emergencyLevel,omitempty"`
	rest string `json:"labels,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

