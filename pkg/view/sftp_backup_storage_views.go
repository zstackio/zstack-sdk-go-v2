// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SftpBackupStorageInventoryView SftpBackupStorage
type SftpBackupStorageInventoryView struct {
	rest string `json:"hostname,omitempty"`
	rest string `json:"username,omitempty"`
	rest int `json:"sshPort,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"description,omitempty"`
	rest int64 `json:"totalCapacity,omitempty"`
	rest int64 `json:"availableCapacity,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []string `json:"attachedZoneUuids,omitempty"`
}

