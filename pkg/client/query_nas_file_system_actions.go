// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryNasFileSystem queries NasFileSystem list
func (cli *ZSClient) QueryNasFileSystem(params *param.QueryParam) ([]view.NasFileSystemInventoryView, error) {
	var resp []view.NasFileSystemInventoryView
	return resp, cli.List("v1/primary-storage/nas", params, &resp)
}
