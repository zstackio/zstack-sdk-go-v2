// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// FlowRouterInventoryView FlowRouter
type FlowRouterInventoryView struct {
	Uuid     string `json:"uuid,omitempty"`
	SystemID int64  `json:"systemID,omitempty"`
	Type     string `json:"type,omitempty"`
}
