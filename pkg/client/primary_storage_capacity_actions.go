// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetPrimaryStorageCapacity 获取PrimaryStorageCapacity详情
func (cli *ZSClient) GetPrimaryStorageCapacity(uuid string) (*view.GetPrimaryStorageCapacityView, error) {
	var resp view.GetPrimaryStorageCapacityView
	if err := cli.Get("v1/primary-storage/capacities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

