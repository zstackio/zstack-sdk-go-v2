// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryGCJob queries GCJob list
func (cli *ZSClient) QueryGCJob(params *param.QueryParam) ([]view.GarbageCollectorInventoryView, error) {
	var resp []view.GarbageCollectorInventoryView
	return resp, cli.List("v1/gc-jobs", params, &resp)
}
