// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MetricRuleTemplateInventoryView MetricRuleTemplate
type MetricRuleTemplateInventoryView struct {
	Name string `json:"name,omitempty"`
	MonitorTemplateUuid string `json:"monitorTemplateUuid,omitempty"`
	ComparisonOperator string `json:"comparisonOperator,omitempty"`
	Period int `json:"period,omitempty"`
	RepeatInterval int `json:"repeatInterval,omitempty"`
	RepeatCount int `json:"repeatCount,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	MetricName string `json:"metricName,omitempty"`
	Threshold float64 `json:"threshold,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	Labels string `json:"labels,omitempty"`
	EnableRecovery bool `json:"enableRecovery,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

