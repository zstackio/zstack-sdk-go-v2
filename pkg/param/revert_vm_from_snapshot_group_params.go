// Copyright (c) ZStack.io, Inc.

package param

// RevertVmFromSnapshotGroupDetailParam RevertVmFromSnapshotGroup detail param
type RevertVmFromSnapshotGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RevertVmFromSnapshotGroupParam RevertVmFromSnapshotGroup request param
type RevertVmFromSnapshotGroupParam struct {
	BaseParam
	Params RevertVmFromSnapshotGroupDetailParam `json:"params"`
}
