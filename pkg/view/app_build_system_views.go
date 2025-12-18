// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AppBuildSystemInventoryView AppBuildSystem
type AppBuildSystemInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"storageType,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"hostname,omitempty"`
	rest string `json:"username,omitempty"`
	rest int `json:"sshPort,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"state,omitempty"`
	rest int64 `json:"totalCapacity,omitempty"`
	rest int64 `json:"availableCapacity,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

