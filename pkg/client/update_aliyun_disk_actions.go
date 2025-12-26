// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAliyunDisk updates AliyunDisk
func (cli *ZSClient) UpdateAliyunDisk(uuid string, params param.UpdateAliyunDiskParam) (*view.UpdateAliyunDiskEventView, error) {
	resp := view.UpdateAliyunDiskEventView{}
	if err := cli.Put("v1/hybrid/aliyun/disk/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
