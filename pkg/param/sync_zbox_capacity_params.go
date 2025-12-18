// Copyright (c) ZStack.io, Inc.

package param

// SyncZBoxCapacityDetailParam SyncZBoxCapacity detail param
type SyncZBoxCapacityDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// SyncZBoxCapacityParam SyncZBoxCapacity request param
type SyncZBoxCapacityParam struct {
	BaseParam
	Params SyncZBoxCapacityDetailParam `json:"params"`
}
