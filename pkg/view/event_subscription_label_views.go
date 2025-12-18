// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EventSubscriptionLabelInventoryView EventSubscriptionLabel
type EventSubscriptionLabelInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Key string `json:"key,omitempty"`
	Operator string `json:"operator,omitempty"`
	Value string `json:"value,omitempty"`
}

