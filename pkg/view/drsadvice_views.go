// Copyright (c) ZStack.io, Inc.

package view

import "time"

// DRSAdviceInventoryView DRSAdvice
type DRSAdviceInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"drsUuid,omitempty"`
	rest string `json:"adviceGroupUuid,omitempty"`
	rest string `json:"vmUuid,omitempty"`
	rest string `json:"vmSourceHostUuid,omitempty"`
	rest string `json:"vmTargetHostUuid,omitempty"`
	rest string `json:"reason,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

