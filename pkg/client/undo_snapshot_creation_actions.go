// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UndoSnapshotCreation 操作UndoSnapshotCreation
func (cli *ZSClient) UndoSnapshotCreation(uuid string, params param.UndoSnapshotCreationParam) (*view.UndoSnapshotCreationEventView, error) {
	resp := view.UndoSnapshotCreationEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

