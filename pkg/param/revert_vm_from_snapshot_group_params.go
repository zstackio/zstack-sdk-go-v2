// Copyright (c) ZStack.io, Inc.

package param

// RevertVmFromSnapshotGroupDetailParam RevertVmFromSnapshotGroup详细参数
type RevertVmFromSnapshotGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// RevertVmFromSnapshotGroupParam RevertVmFromSnapshotGroup请求参数
type RevertVmFromSnapshotGroupParam struct {
	BaseParam
	Params RevertVmFromSnapshotGroupDetailParam `json:"params"` // 详细参数
}

