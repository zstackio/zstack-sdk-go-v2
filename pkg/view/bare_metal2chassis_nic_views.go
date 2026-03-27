// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BareMetal2ChassisNicInventoryView BareMetal2ChassisNic
type BareMetal2ChassisNicInventoryView struct {
	BaseInfoView
	BaseTimeView
	ChassisUuid string `json:"chassisUuid,omitempty"`
	Mac string `json:"mac,omitempty"`
	NicName string `json:"nicName,omitempty"`
	Speed string `json:"speed,omitempty"`
	IsProvisionNic bool `json:"isProvisionNic,omitempty"`
}

