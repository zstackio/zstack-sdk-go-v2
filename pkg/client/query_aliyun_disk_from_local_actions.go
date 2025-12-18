// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAliyunDiskFromLocal queries AliyunDiskFromLocal list
func (cli *ZSClient) QueryAliyunDiskFromLocal(params param.QueryParam) ([]view.AliyunDiskInventoryView, error) {
	var resp []view.AliyunDiskInventoryView
	return resp, cli.List("v1/hybrid/aliyun/disk", &params, &resp)
}
