// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ModelCenterServiceInventoryView ModelCenterService
type ModelCenterServiceInventoryView struct {
	BaseInfoView
	BaseTimeView
	ServiceStatuses []ServiceStatusView `json:"serviceStatuses,omitempty"`
	Zdfs ZdfsServiceView `json:"zdfs,omitempty"`
}

