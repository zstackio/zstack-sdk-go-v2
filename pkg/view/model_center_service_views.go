// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ModelCenterServiceInventoryView ModelCenterService
type ModelCenterServiceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ServiceStatuses []ServiceStatusView `json:"serviceStatuses,omitempty"`
	Zdfs ZdfsServiceView `json:"zdfs,omitempty"`
}

