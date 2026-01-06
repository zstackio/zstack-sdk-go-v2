// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MetricDataHttpReceiverInventoryView MetricDataHttpReceiver
type MetricDataHttpReceiverInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	State string `json:"state,omitempty"`
}

// DeleteMetricDataHttpReceiverEventView DeleteMetricDataHttpReceiverEvent
type DeleteMetricDataHttpReceiverEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateMetricDataHttpReceiverEventView CreateMetricDataHttpReceiverEvent
type CreateMetricDataHttpReceiverEventView struct {
	Inventory MetricDataHttpReceiverInventoryView `json:"inventory,omitempty"`
}

