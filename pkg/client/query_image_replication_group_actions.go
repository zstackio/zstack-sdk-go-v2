// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryImageReplicationGroup queries ImageReplicationGroup list
func (cli *ZSClient) QueryImageReplicationGroup(params *param.QueryParam) ([]view.ImageReplicationGroupInventoryView, error) {
	var resp []view.ImageReplicationGroupInventoryView
	return resp, cli.List("v1/image-replication-groups", params, &resp)
}
