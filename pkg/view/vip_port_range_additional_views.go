// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VipPortRangeInventoryView VipPortRange
type VipPortRangeInventoryView struct {
	Uuid      string   `json:"uuid,omitempty"`
	Protocol  string   `json:"protocol,omitempty"`
	UsedPorts []string `json:"usedPorts,omitempty"`
}
