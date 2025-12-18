// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VipInventoryView Vip
type VipInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"ip,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"gateway,omitempty"`
	rest string `json:"netmask,omitempty"`
	rest int `json:"prefixLen,omitempty"`
	rest string `json:"serviceProvider,omitempty"`
	rest []string `json:"peerL3NetworkUuids,omitempty"`
	rest []VipNetworkServicesRefInventoryView `json:"servicesRefs,omitempty"`
	rest string `json:"useFor,omitempty"`
	rest bool `json:"system,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

