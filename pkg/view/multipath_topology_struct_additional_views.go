// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MultipathTopologyStructView MultipathTopologyStruct
type MultipathTopologyStructView struct {
	DiskUuid *string `json:"diskUuid,omitempty"`
	Devices []DeviceTOView `json:"devices,omitempty"`
}

