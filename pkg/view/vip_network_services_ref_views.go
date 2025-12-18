// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VipNetworkServicesRefInventoryView VipNetworkServicesRef
type VipNetworkServicesRefInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"serviceType,omitempty"`
	rest string `json:"vipUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

