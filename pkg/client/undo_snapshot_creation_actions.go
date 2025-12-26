// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UndoSnapshotCreation operates on UndoSnapshotCreation
func (cli *ZSClient) UndoSnapshotCreation(uuid string, params param.UndoSnapshotCreationParam) (*view.UndoSnapshotCreationEventView, error) {
	resp := view.UndoSnapshotCreationEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
