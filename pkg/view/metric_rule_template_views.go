// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MetricRuleTemplateInventoryView MetricRuleTemplate
type MetricRuleTemplateInventoryView struct {
	rest string `json:"name,omitempty"`
	rest string `json:"monitorTemplateUuid,omitempty"`
	rest string `json:"comparisonOperator,omitempty"`
	rest int `json:"period,omitempty"`
	rest int `json:"repeatInterval,omitempty"`
	rest int `json:"repeatCount,omitempty"`
	rest string `json:"namespace,omitempty"`
	rest string `json:"metricName,omitempty"`
	rest float64 `json:"threshold,omitempty"`
	rest string `json:"emergencyLevel,omitempty"`
	rest string `json:"labels,omitempty"`
	rest bool `json:"enableRecovery,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"uuid,omitempty"`
}

