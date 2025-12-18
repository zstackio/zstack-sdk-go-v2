// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BaremetalInstanceInventoryView BaremetalInstance
type BaremetalInstanceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	PxeServerUuid string `json:"pxeServerUuid,omitempty"`
	ChassisUuid string `json:"chassisUuid,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
	TemplateUuid string `json:"templateUuid,omitempty"`
	Platform string `json:"platform,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Username string `json:"username,omitempty"`
	Port int `json:"port,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	BmNics []BaremetalNicInventoryView `json:"bmNics,omitempty"`
}

