// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVolumeIoThreadPin 获取VolumeIoThreadPin详情
func (cli *ZSClient) GetVolumeIoThreadPin(uuid string) (*view.GetVolumeIoThreadPinView, error) {
	var resp view.GetVolumeIoThreadPinView
	if err := cli.Get("v1/volumes/{uuid}/io-thread-pin", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

