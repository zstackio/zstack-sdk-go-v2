// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddAliyunPanguPartition 操作AddAliyunPanguPartition
func (cli *ZSClient) AddAliyunPanguPartition(params param.AddAliyunPanguPartitionParam) (*view.AddAliyunPanguPartitionEventView, error) {
	resp := view.AddAliyunPanguPartitionEventView{}
	if err := cli.Post("v1/aliyun/pangu", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

