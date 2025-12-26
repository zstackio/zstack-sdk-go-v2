// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAliyunDiskFromLocal queries AliyunDiskFromLocal list
func (cli *ZSClient) QueryAliyunDiskFromLocal(params *param.QueryParam) ([]view.AliyunDiskInventoryView, error) {
	var resp []view.AliyunDiskInventoryView
	return resp, cli.List("v1/hybrid/aliyun/disk", params, &resp)
}
