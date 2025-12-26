// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddAliyunPanguPartition adds AliyunPanguPartition
func (cli *ZSClient) AddAliyunPanguPartition(params param.AddAliyunPanguPartitionParam) (*view.AddAliyunPanguPartitionEventView, error) {
	resp := view.AddAliyunPanguPartitionEventView{}
	if err := cli.Post("v1/aliyun/pangu", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
