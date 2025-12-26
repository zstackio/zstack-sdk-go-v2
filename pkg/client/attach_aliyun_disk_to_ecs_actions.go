// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachAliyunDiskToEcs operates on AliyunDiskToEcs
func (cli *ZSClient) AttachAliyunDiskToEcs(params param.AttachAliyunDiskToEcsParam) (*view.AttachAliyunDiskToEcsEventView, error) {
	resp := view.AttachAliyunDiskToEcsEventView{}
	if err := cli.Post("v1/hybrid/aliyun/disk/{diskUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
