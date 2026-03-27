// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LongJobProgressDetailView LongJobProgressDetail
type LongJobProgressDetailView struct {
	Percent int `json:"percent,omitempty"`
	Stage string `json:"stage,omitempty"`
	State string `json:"state,omitempty"`
	StateReason string `json:"stateReason,omitempty"`
	ProcessedBytes int64 `json:"processedBytes,omitempty"`
	TotalBytes int64 `json:"totalBytes,omitempty"`
	ProcessedItems int64 `json:"processedItems,omitempty"`
	TotalItems int64 `json:"totalItems,omitempty"`
	SpeedBytesPerSecond int64 `json:"speedBytesPerSecond,omitempty"`
	EstimatedRemainingSeconds int64 `json:"estimatedRemainingSeconds,omitempty"`
	Extra map[string]interface{} `json:"extra,omitempty"`
}

