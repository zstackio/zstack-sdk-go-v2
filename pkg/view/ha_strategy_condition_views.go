// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HaStrategyConditionInventoryView HaStrategyCondition
type HaStrategyConditionInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"fencerName,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

