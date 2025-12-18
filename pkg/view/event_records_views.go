// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EventRecordsInventoryView EventRecords
type EventRecordsInventoryView struct {
	Id int64 `json:"id,omitempty"`
	CreateTime int64 `json:"createTime,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Name string `json:"name,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	ResourceId string `json:"resourceId,omitempty"`
	DataUuid string `json:"dataUuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	SubscriptionUuid string `json:"subscriptionUuid,omitempty"`
	ReadStatus bool `json:"readStatus,omitempty"`
	Labels string `json:"labels,omitempty"`
}

