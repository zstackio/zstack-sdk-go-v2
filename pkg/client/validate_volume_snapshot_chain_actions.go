// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ValidateVolumeSnapshotChain 操作ValidateVolumeSnapshotChain
func (cli *ZSClient) ValidateVolumeSnapshotChain(uuid string, params param.ValidateVolumeSnapshotChainParam) (*view.ValidateVolumeSnapshotChainEventView, error) {
	resp := view.ValidateVolumeSnapshotChainEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

