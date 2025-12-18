// Copyright (c) ZStack.io, Inc.

package param

// SetVmClockTrackDetailParam SetVmClockTrack详细参数
type SetVmClockTrackDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"track" validate:"required"` // 必填
	rest bool `json:"syncAfterVMResume,omitempty"`
	rest int `json:"intervalInSeconds,omitempty"`
}

// SetVmClockTrackParam SetVmClockTrack请求参数
type SetVmClockTrackParam struct {
	BaseParam
	Params SetVmClockTrackDetailParam `json:"params"` // 详细参数
}

