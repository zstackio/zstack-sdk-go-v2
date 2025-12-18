// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetImageQga 获取ImageQga详情
func (cli *ZSClient) GetImageQga(uuid string) (*view.GetImageQgaView, error) {
	var resp view.GetImageQgaView
	if err := cli.Get("v1/images/{uuid}/qga", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

