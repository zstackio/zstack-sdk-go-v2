// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VipNetworkServicesRefInventoryView VipNetworkServicesRef
type VipNetworkServicesRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

