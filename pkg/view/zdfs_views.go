// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ZdfsInventoryView Zdfs
type ZdfsInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"hostName,omitempty"`
	rest int `json:"sshPort,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest ZdfsStorageInventoryView `json:"storage,omitempty"`
}

