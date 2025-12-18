// Copyright (c) ZStack.io, Inc.

package view

import "time"

// KvmHypervisorInfoInventoryView KvmHypervisorInfo
type KvmHypervisorInfoInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"hypervisor,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"matchState,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

