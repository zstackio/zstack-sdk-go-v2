// Copyright (c) ZStack.io, Inc.

package param

// SetVmClockTrackDetailParam SetVmClockTrack detail param
type SetVmClockTrackDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Track string `json:"track" validate:"required"`
	SyncAfterVMResume bool `json:"syncAfterVMResume,omitempty"`
	IntervalInSeconds int `json:"intervalInSeconds,omitempty"`
}

// SetVmClockTrackParam SetVmClockTrack request param
type SetVmClockTrackParam struct {
	BaseParam
	Params SetVmClockTrackDetailParam `json:"params"`
}
