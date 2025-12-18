// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddSharedMountPointPrimaryStorage 操作AddSharedMountPointPrimaryStorage
func (cli *ZSClient) AddSharedMountPointPrimaryStorage(params param.AddSharedMountPointPrimaryStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/smp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

