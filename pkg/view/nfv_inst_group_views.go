// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NfvInstGroupInventoryView NfvInstGroup
type NfvInstGroupInventoryView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Description *string `json:"description,omitempty"`
	NfvInstOfferingUuid *string `json:"nfvInstOfferingUuid,omitempty"`
	InstType string `json:"instType,omitempty"`
	FuncType string `json:"funcType,omitempty"`
	ConfigVersion int `json:"configVersion,omitempty"`
	NetOsDistro *string `json:"netOsDistro,omitempty"`
	BaseOsDistro *string `json:"baseOsDistro,omitempty"`
	Status string `json:"status,omitempty"`
	StatusDetail *string `json:"statusDetail,omitempty"`
	OperationMode string `json:"operationMode,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	Instances []NfvInstInventoryView `json:"instances,omitempty"`
	Monitors []NfvInstGroupMonitorIpInventoryView `json:"monitors,omitempty"`
	Services []NfvInstGroupNetworkServiceRefInventoryView `json:"services,omitempty"`
	ConfigTasks []NfvInstGroupConfigTaskInventoryView `json:"configTasks,omitempty"`
	L3Networks []NfvInstGroupL3NetworkRefInventoryView `json:"l3Networks,omitempty"`
	VipUuid *string `json:"vipUuid,omitempty"`
	Ipv6VipUuid *string `json:"ipv6VipUuid,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	PrimaryStoragePoolUuid *string `json:"primaryStoragePoolUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
}

// CreateNfvInstGroupEventView CreateNfvInstGroupEvent
type CreateNfvInstGroupEventView struct {
	Inventory NfvInstGroupInventoryView `json:"inventory,omitempty"`
}

// SyncNfvInstGroupEventView SyncNfvInstGroupEvent
type SyncNfvInstGroupEventView struct {
	Inventory NfvInstGroupInventoryView `json:"inventory,omitempty"`
}

// UpdateNfvInstGroupEventView UpdateNfvInstGroupEvent
type UpdateNfvInstGroupEventView struct {
	Inventory NfvInstGroupInventoryView `json:"inventory,omitempty"`
}

// DetachNfvInstFromGroupEventView DetachNfvInstFromGroupEvent
type DetachNfvInstFromGroupEventView struct {
	Inventory NfvInstGroupInventoryView `json:"inventory,omitempty"`
}

// ChangeNfvInstGroupOperationModeEventView ChangeNfvInstGroupOperationModeEvent
type ChangeNfvInstGroupOperationModeEventView struct {
	Inventory NfvInstGroupInventoryView `json:"inventory,omitempty"`
}

// AttachNfvInstToGroupEventView AttachNfvInstToGroupEvent
type AttachNfvInstToGroupEventView struct {
	Inventory NfvInstGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteNfvInstGroupEventView DeleteNfvInstGroupEvent
type DeleteNfvInstGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryNfvInstGroupView QueryNfvInstGroup
type QueryNfvInstGroupView struct {
	Inventories []NfvInstGroupInventoryView `json:"inventories,omitempty"`
}

