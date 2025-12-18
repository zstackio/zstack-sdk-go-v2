// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAliyunPanguPartition 查询AliyunPanguPartition列表
func (cli *ZSClient) QueryAliyunPanguPartition(params param.QueryParam) ([]view.QueryAliyunPanguPartitionView, error) {
	var resp []view.QueryAliyunPanguPartitionView
	return resp, cli.List("v1/aliyun/pangu", &params, &resp)
}

