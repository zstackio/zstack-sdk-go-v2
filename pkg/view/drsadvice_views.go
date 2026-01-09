// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// DRSAdviceInventoryView DRSAdvice
type DRSAdviceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	DrsUuid *string `json:"drsUuid,omitempty"`
	AdviceGroupUuid *string `json:"adviceGroupUuid,omitempty"`
	VmUuid *string `json:"vmUuid,omitempty"`
	VmSourceHostUuid *string `json:"vmSourceHostUuid,omitempty"`
	VmTargetHostUuid *string `json:"vmTargetHostUuid,omitempty"`
	Reason *string `json:"reason,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// ApplyDRSAdviceEventView ApplyDRSAdviceEvent
type ApplyDRSAdviceEventView struct {
	VmMigrationActivityUuid *string `json:"vmMigrationActivityUuid,omitempty"`
	Success bool `json:"success,omitempty"`
}

// QueryDRSAdviceView QueryDRSAdvice
type QueryDRSAdviceView struct {
	Inventories []DRSAdviceInventoryView `json:"inventories,omitempty"`
}

