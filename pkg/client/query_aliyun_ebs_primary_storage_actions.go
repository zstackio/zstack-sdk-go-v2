// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAliyunEbsPrimaryStorage queries AliyunEbsPrimaryStorage list
func (cli *ZSClient) QueryAliyunEbsPrimaryStorage(params param.QueryParam) ([]view.PrimaryStorageInventoryView, error) {
	var resp []view.PrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage/aliyun/ebs", &params, &resp)
}
