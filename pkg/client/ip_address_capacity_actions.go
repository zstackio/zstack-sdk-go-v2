// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetIpAddressCapacity 获取IpAddressCapacity详情
func (cli *ZSClient) GetIpAddressCapacity(uuid string) (*view.GetIpAddressCapacityView, error) {
	var resp view.GetIpAddressCapacityView
	if err := cli.Get("v1/ip-capacity", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

