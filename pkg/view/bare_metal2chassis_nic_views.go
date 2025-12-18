// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BareMetal2ChassisNicInventoryView BareMetal2ChassisNic
type BareMetal2ChassisNicInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"chassisUuid,omitempty"`
	rest string `json:"mac,omitempty"`
	rest string `json:"nicName,omitempty"`
	rest string `json:"speed,omitempty"`
	rest bool `json:"isProvisionNic,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

