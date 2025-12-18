// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MetricTemplateInventoryView MetricTemplate
type MetricTemplateInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"receiverUuid,omitempty"`
	rest string `json:"template,omitempty"`
	rest string `json:"namespace,omitempty"`
	rest string `json:"metricName,omitempty"`
	rest string `json:"labelsJsonStr,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

