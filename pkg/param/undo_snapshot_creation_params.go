// Copyright (c) ZStack.io, Inc.

package param

// UndoSnapshotCreationDetailParam UndoSnapshotCreation详细参数
type UndoSnapshotCreationDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"snapShotUuid" validate:"required"` // 必填
}

// UndoSnapshotCreationParam UndoSnapshotCreation请求参数
type UndoSnapshotCreationParam struct {
	BaseParam
	Params UndoSnapshotCreationDetailParam `json:"params"` // 详细参数
}

