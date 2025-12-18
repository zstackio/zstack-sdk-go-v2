// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAliyunDisk 更新AliyunDisk
func (cli *ZSClient) UpdateAliyunDisk(uuid string, params param.UpdateAliyunDiskParam) (*view.UpdateAliyunDiskEventView, error) {
	resp := view.UpdateAliyunDiskEventView{}
	if err := cli.Put("v1/hybrid/aliyun/disk/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

