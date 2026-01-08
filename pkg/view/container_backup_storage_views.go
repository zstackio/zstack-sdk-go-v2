// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ContainerBackupStorageInventoryView ContainerBackupStorage
type ContainerBackupStorageInventoryView struct {
	BaseInfoView
	BaseTimeView
	EndpointUuid      string   `json:"endpointUuid,omitempty"`
	Id                int64    `json:"id,omitempty"`
	Url               string   `json:"url,omitempty"`
	TotalCapacity     int64    `json:"totalCapacity,omitempty"`
	AvailableCapacity int64    `json:"availableCapacity,omitempty"`
	Type              string   `json:"type,omitempty"`
	State             string   `json:"state,omitempty"`
	Status            string   `json:"status,omitempty"`
	AttachedZoneUuids []string `json:"attachedZoneUuids,omitempty"`
}
