// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ExternalPrimaryStorageInventoryView ExternalPrimaryStorage
type ExternalPrimaryStorageInventoryView struct {
	Identity *string `json:"identity,omitempty"`
	Config interface{} `json:"config,omitempty"`
	AddonInfo interface{} `json:"addonInfo,omitempty"`
	OutputProtocols []string `json:"outputProtocols,omitempty"`
	DefaultProtocol *string `json:"defaultProtocol,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Url *string `json:"url,omitempty"`
	Description *string `json:"description,omitempty"`
	TotalCapacity *int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity *int64 `json:"availableCapacity,omitempty"`
	TotalPhysicalCapacity *int64 `json:"totalPhysicalCapacity,omitempty"`
	AvailablePhysicalCapacity *int64 `json:"availablePhysicalCapacity,omitempty"`
	SystemUsedCapacity *int64 `json:"systemUsedCapacity,omitempty"`
	Type *string `json:"type,omitempty"`
	State *string `json:"state,omitempty"`
	Status *string `json:"status,omitempty"`
	MountPath *string `json:"mountPath,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
}

// UpdateExternalPrimaryStorageEventView UpdateExternalPrimaryStorageEvent
type UpdateExternalPrimaryStorageEventView struct {
	Inventory ExternalPrimaryStorageInventoryView `json:"inventory,omitempty"`
}

