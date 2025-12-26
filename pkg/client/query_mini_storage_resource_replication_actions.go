// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryMiniStorageResourceReplication queries MiniStorageResourceReplication list
func (cli *ZSClient) QueryMiniStorageResourceReplication(params *param.QueryParam) ([]view.MiniStorageResourceReplicationInventoryView, error) {
	var resp []view.MiniStorageResourceReplicationInventoryView
	return resp, cli.List("v1/primary-storage/mini/replications", params, &resp)
}
