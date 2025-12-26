// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RevertVmFromSnapshotGroup operates on RevertVmFromSnapshotGroup
func (cli *ZSClient) RevertVmFromSnapshotGroup(uuid string, params param.RevertVmFromSnapshotGroupParam) (*view.RevertVmFromSnapshotGroupEventView, error) {
	resp := view.RevertVmFromSnapshotGroupEventView{}
	if err := cli.Put("v1/volume-snapshots/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
