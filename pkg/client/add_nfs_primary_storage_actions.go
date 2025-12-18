// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddNfsPrimaryStorage 操作AddNfsPrimaryStorage
func (cli *ZSClient) AddNfsPrimaryStorage(params param.AddNfsPrimaryStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/nfs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

