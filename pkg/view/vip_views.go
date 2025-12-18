// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VipInventoryView Vip
type VipInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	Ip string `json:"ip,omitempty"`
	State string `json:"state,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	PrefixLen int `json:"prefixLen,omitempty"`
	ServiceProvider string `json:"serviceProvider,omitempty"`
	PeerL3NetworkUuids []string `json:"peerL3NetworkUuids,omitempty"`
	ServicesRefs []VipNetworkServicesRefInventoryView `json:"servicesRefs,omitempty"`
	UseFor string `json:"useFor,omitempty"`
	System bool `json:"system,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

