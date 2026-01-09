// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SanSecSecurityMachineInventoryView SanSecSecurityMachine
type SanSecSecurityMachineInventoryView struct {
	Port *int `json:"port,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	Name string `json:"name,omitempty"`
	SecretResourcePoolUuid *string `json:"secretResourcePoolUuid,omitempty"`
	Description *string `json:"description,omitempty"`
	ManagementIp *string `json:"managementIp,omitempty"`
	Type *string `json:"type,omitempty"`
	Model *string `json:"model,omitempty"`
	State *string `json:"state,omitempty"`
	Status *string `json:"status,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

