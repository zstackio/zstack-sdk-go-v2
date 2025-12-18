// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AlarmInventoryView Alarm
type AlarmInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ComparisonOperator string `json:"comparisonOperator,omitempty"`
	Period int `json:"period,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	MetricName string `json:"metricName,omitempty"`
	Threshold float64 `json:"threshold,omitempty"`
	RepeatInterval int `json:"repeatInterval,omitempty"`
	RepeatCount int `json:"repeatCount,omitempty"`
	Status string `json:"status,omitempty"`
	State string `json:"state,omitempty"`
	EnableRecovery bool `json:"enableRecovery,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Labels []AlarmLabelInventoryView `json:"labels,omitempty"`
	Actions []AlarmActionInventoryView `json:"actions,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
}

