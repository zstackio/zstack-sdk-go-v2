// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VmPriorityConfigInventoryView VmPriorityConfig
type VmPriorityConfigInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	Level string `json:"level,omitempty"`
	CpuShares int `json:"cpuShares,omitempty"`
	OomScoreAdj int `json:"oomScoreAdj,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

