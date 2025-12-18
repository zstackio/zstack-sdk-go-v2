// Copyright (c) ZStack.io, Inc.

package param

// SyncVmClockDetailParam SyncVmClock detail param
type SyncVmClockDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// SyncVmClockParam SyncVmClock request param
type SyncVmClockParam struct {
	BaseParam
	Params SyncVmClockDetailParam `json:"params"`
}
