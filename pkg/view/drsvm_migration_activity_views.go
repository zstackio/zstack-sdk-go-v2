// Copyright (c) ZStack.io, Inc.

package view

import "time"

// DRSVmMigrationActivityInventoryView DRSVmMigrationActivity
type DRSVmMigrationActivityInventoryView struct {
	rest string `json:"drsUuid,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"vmUuid,omitempty"`
	rest string `json:"vmSourceHostUuid,omitempty"`
	rest string `json:"vmTargetHostUuid,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"result,omitempty"`
	rest string `json:"reason,omitempty"`
	rest string `json:"adviceUuid,omitempty"`
	rest string `json:"cause,omitempty"`
	rest time.Time `json:"endDate,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

