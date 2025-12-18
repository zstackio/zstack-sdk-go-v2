// Copyright (c) ZStack.io, Inc.

package param

// UndoSnapshotCreationDetailParam UndoSnapshotCreation detail param
type UndoSnapshotCreationDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SnapShotUuid string `json:"snapShotUuid" validate:"required"`
}

// UndoSnapshotCreationParam UndoSnapshotCreation request param
type UndoSnapshotCreationParam struct {
	BaseParam
	Params UndoSnapshotCreationDetailParam `json:"params"`
}
