// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2UsageInventoryView BareMetal2Usage
type BareMetal2UsageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	BareMetal2ChassisOfferingUuid string `json:"bareMetal2ChassisOfferingUuid,omitempty"`
	VmUuid string `json:"vmUuid,omitempty"`
	VmName string `json:"vmName,omitempty"`
	State string `json:"state,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	DateInLong int64 `json:"dateInLong,omitempty"`
}

