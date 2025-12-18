// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMiniStorageResourceReplication queries MiniStorageResourceReplication list
func (cli *ZSClient) QueryMiniStorageResourceReplication(params param.QueryParam) ([]view.MiniStorageResourceReplicationInventoryView, error) {
	var resp []view.MiniStorageResourceReplicationInventoryView
	return resp, cli.List("v1/primary-storage/mini/replications", &params, &resp)
}
