// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DetachAliyunDiskFromEcs operates on AliyunDiskFromEcs
func (cli *ZSClient) DetachAliyunDiskFromEcs(params param.DetachAliyunDiskFromEcsParam) (*view.DetachAliyunDiskFromEcsEventView, error) {
	resp := view.DetachAliyunDiskFromEcsEventView{}
	if err := cli.Post("v1/hybrid/aliyun/disk/{uuid}/detach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
