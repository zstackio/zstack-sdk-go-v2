// Copyright (c) ZStack.io, Inc.

package view

// ExportNbdVolumesEventView ExportNbdVolumesEvent
type ExportNbdVolumesEventView struct {
	VolumeInfos []interface{} `json:"volumeInfos,omitempty"`
	Success bool `json:"success,omitempty"`
}

