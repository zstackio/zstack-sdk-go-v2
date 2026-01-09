// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ActiveAlarmTemplateInventoryView ActiveAlarmTemplate
type ActiveAlarmTemplateInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AlarmName *string `json:"alarmName,omitempty"`
	ComparisonOperator string `json:"comparisonOperator,omitempty"`
	Period int `json:"period,omitempty"`
	RepeatInterval int `json:"repeatInterval,omitempty"`
	RepeatCount int `json:"repeatCount,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	MetricName *string `json:"metricName,omitempty"`
	Threshold float64 `json:"threshold,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	Labels *string `json:"labels,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// QueryActiveAlarmTemplateView QueryActiveAlarmTemplate
type QueryActiveAlarmTemplateView struct {
	Inventories []ActiveAlarmTemplateInventoryView `json:"inventories,omitempty"`
}

// UpdateActiveAlarmTemplateEventView UpdateActiveAlarmTemplateEvent
type UpdateActiveAlarmTemplateEventView struct {
	Inventory ActiveAlarmTemplateInventoryView `json:"inventory,omitempty"`
}

