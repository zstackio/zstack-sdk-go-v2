// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RefreshNvmeTarget operates on RefreshNvmeTarget
func (cli *ZSClient) RefreshNvmeTarget(params param.RefreshNvmeTargetParam) (*view.RefreshNvmeTargetEventView, error) {
	resp := view.RefreshNvmeTargetEventView{}
	if err := cli.Post("v1/storage-devices/nvme/controllers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
