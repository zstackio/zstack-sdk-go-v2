// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VCenterClusterInventoryView VCenterCluster
type VCenterClusterInventoryView struct {
	VCenterUuid *string `json:"vCenterUuid,omitempty"`
	Morval *string `json:"morval,omitempty"`
	DataCenterUuid *string `json:"dataCenterUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Description *string `json:"description,omitempty"`
	State *string `json:"state,omitempty"`
	HypervisorType *string `json:"hypervisorType,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	Type *string `json:"type,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
}

// QueryVCenterClusterView QueryVCenterCluster
type QueryVCenterClusterView struct {
	Inventories []VCenterClusterInventoryView `json:"inventories,omitempty"`
}

