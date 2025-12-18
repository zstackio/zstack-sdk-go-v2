// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RefreshNvmeTarget 操作RefreshNvmeTarget
func (cli *ZSClient) RefreshNvmeTarget(params param.RefreshNvmeTargetParam) (*view.RefreshNvmeTargetEventView, error) {
	resp := view.RefreshNvmeTargetEventView{}
	if err := cli.Post("v1/storage-devices/nvme/controllers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

