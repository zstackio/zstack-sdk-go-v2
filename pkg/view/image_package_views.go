// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ImagePackageInventoryView ImagePackage
type ImagePackageInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"vmUuid,omitempty"`
	rest string `json:"backupStorageUuid,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"exportUrl,omitempty"`
	rest string `json:"md5Sum,omitempty"`
	rest string `json:"format,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

