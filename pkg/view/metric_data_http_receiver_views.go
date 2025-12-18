// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MetricDataHttpReceiverInventoryView MetricDataHttpReceiver
type MetricDataHttpReceiverInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"state,omitempty"`
}

