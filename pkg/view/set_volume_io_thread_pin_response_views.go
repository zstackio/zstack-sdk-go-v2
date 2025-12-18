// Copyright (c) ZStack.io, Inc.

package view

// SetVolumeIoThreadPinEventView SetVolumeIoThreadPinEvent
type SetVolumeIoThreadPinEventView struct {
	VolumeUuid string `json:"volumeUuid,omitempty"`
	IoThreadId int `json:"ioThreadId,omitempty"`
	Pin string `json:"pin,omitempty"`
}

