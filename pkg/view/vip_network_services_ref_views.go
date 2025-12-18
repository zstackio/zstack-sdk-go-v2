// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VipNetworkServicesRefInventoryView VipNetworkServicesRef
type VipNetworkServicesRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

