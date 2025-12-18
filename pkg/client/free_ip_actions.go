// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetFreeIp 获取FreeIp详情
func (cli *ZSClient) GetFreeIp(uuid string) (*view.GetFreeIpView, error) {
	var resp view.GetFreeIpView
	if err := cli.Get("v1/l3-networks/ip/free", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

