// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LongJobProgressDetailView LongJobProgressDetail
type LongJobProgressDetailView struct {
	Percent int `json:"percent,omitempty"`
	Stage string `json:"stage,omitempty"`
	State string `json:"state,omitempty"`
	StateReason string `json:"stateReason,omitempty"`
	Processed int64 `json:"processed,omitempty"`
	Total int64 `json:"total,omitempty"`
	ProcessedItems int64 `json:"processedItems,omitempty"`
	TotalItems int64 `json:"totalItems,omitempty"`
	Speed int64 `json:"speed,omitempty"`
	Unit string `json:"unit,omitempty"`
	EstimatedRemainingSeconds int64 `json:"estimatedRemainingSeconds,omitempty"`
	Extra map[string]interface{} `json:"extra,omitempty"`
}

