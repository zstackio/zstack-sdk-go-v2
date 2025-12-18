// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddBlockPrimaryStorage adds BlockPrimaryStorage
func (cli *ZSClient) AddBlockPrimaryStorage(params param.AddBlockPrimaryStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/block", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
