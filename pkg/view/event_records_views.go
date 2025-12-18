// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EventRecordsInventoryView EventRecords
type EventRecordsInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest int64 `json:"createTime,omitempty"`
	rest string `json:"namespace,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"emergencyLevel,omitempty"`
	rest string `json:"resourceId,omitempty"`
	rest string `json:"dataUuid,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"subscriptionUuid,omitempty"`
	rest bool `json:"readStatus,omitempty"`
	rest string `json:"labels,omitempty"`
}

