// Copyright (c) ZStack.io, Inc.

package view

import "time"

// KVMIsoTOView KVMIsoTO
type KVMIsoTOView struct {
	rest string `json:"pathInCache,omitempty"`
	rest string `json:"installUrl,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest int64 `json:"actualSize,omitempty"`
	rest string `json:"md5Sum,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"mediaType,omitempty"`
	rest string `json:"guestOsType,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"platform,omitempty"`
	rest string `json:"architecture,omitempty"`
	rest string `json:"format,omitempty"`
	rest bool `json:"system,omitempty"`
	rest bool `json:"virtio,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []ImageBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
	rest []SystemTagInventoryView `json:"systemTags,omitempty"`
}

