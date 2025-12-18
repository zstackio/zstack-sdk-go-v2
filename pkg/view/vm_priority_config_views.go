// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VmPriorityConfigInventoryView VmPriorityConfig
type VmPriorityConfigInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"level,omitempty"`
	rest int `json:"cpuShares,omitempty"`
	rest int `json:"oomScoreAdj,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

