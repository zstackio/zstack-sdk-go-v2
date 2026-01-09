// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// GarbageCollectorInventoryView GarbageCollector
type GarbageCollectorInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	RunnerClass *string `json:"runnerClass,omitempty"`
	Context *string `json:"context,omitempty"`
	Status *string `json:"status,omitempty"`
	ManagementNodeUuid *string `json:"managementNodeUuid,omitempty"`
	Type *string `json:"type,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// QueryGCJobView QueryGCJob
type QueryGCJobView struct {
	Inventories []GarbageCollectorInventoryView `json:"inventories,omitempty"`
}

