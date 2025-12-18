// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmInstanceHaLevel 获取VmInstanceHaLevel详情
func (cli *ZSClient) GetVmInstanceHaLevel(uuid string) (*view.GetVmInstanceHaLevelView, error) {
	var resp view.GetVmInstanceHaLevelView
	if err := cli.Get("v1/vm-instances/{uuid}/ha-levels", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

