// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAliyunPanguPartition updates AliyunPanguPartition
func (cli *ZSClient) UpdateAliyunPanguPartition(uuid string, params param.UpdateAliyunPanguPartitionParam) (*view.UpdateAliyunPanguPartitionEventView, error) {
	resp := view.UpdateAliyunPanguPartitionEventView{}
	if err := cli.Put("v1/aliyun/pangu/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
