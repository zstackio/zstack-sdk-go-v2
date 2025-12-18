// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMiniStorageResourceReplication 查询MiniStorageResourceReplication列表
func (cli *ZSClient) QueryMiniStorageResourceReplication(params param.QueryParam) ([]view.QueryMiniStorageResourceReplicationView, error) {
	var resp []view.QueryMiniStorageResourceReplicationView
	return resp, cli.List("v1/primary-storage/mini/replications", &params, &resp)
}

