// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SNSSnmpPlatformInventoryView SNSSnmpPlatform
type SNSSnmpPlatformInventoryView struct {
	rest string `json:"snmpAddress,omitempty"`
	rest int `json:"snmpPort,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

