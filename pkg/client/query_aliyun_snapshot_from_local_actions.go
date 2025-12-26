// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAliyunSnapshotFromLocal queries AliyunSnapshotFromLocal list
func (cli *ZSClient) QueryAliyunSnapshotFromLocal(params *param.QueryParam) ([]view.AliyunSnapshotInventoryView, error) {
	var resp []view.AliyunSnapshotInventoryView
	return resp, cli.List("v1/hybrid/aliyun/snapshot", params, &resp)
}
