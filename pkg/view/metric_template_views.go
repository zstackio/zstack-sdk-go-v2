// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MetricTemplateInventoryView MetricTemplate
type MetricTemplateInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ReceiverUuid string `json:"receiverUuid,omitempty"`
	Template string `json:"template,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	MetricName string `json:"metricName,omitempty"`
	LabelsJsonStr string `json:"labelsJsonStr,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

