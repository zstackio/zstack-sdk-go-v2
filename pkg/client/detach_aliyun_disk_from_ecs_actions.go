// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachAliyunDiskFromEcs operates on AliyunDiskFromEcs
func (cli *ZSClient) DetachAliyunDiskFromEcs(params param.DetachAliyunDiskFromEcsParam) (*view.DetachAliyunDiskFromEcsEventView, error) {
	resp := view.DetachAliyunDiskFromEcsEventView{}
	if err := cli.Post("v1/hybrid/aliyun/disk/{uuid}/detach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
