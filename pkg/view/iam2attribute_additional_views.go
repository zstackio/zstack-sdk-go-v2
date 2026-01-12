// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2AttributeInventoryView IAM2Attribute
type IAM2AttributeInventoryView struct {
	BaseInfoView
	BaseTimeView
	Value *string `json:"value,omitempty"`
	Type string `json:"type,omitempty"`
}

