// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VCenterInventoryView VCenter
type VCenterInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DomainName string `json:"domainName,omitempty"`
	Port int `json:"port,omitempty"`
	UserName string `json:"userName,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Version string `json:"version,omitempty"`
	Https bool `json:"https,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

