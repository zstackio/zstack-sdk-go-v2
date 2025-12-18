// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetBareMetal2ProvisionNetworkIpAddressCapacity 获取BareMetal2ProvisionNetworkIpAddressCapacity详情
func (cli *ZSClient) GetBareMetal2ProvisionNetworkIpAddressCapacity(uuid string) (*view.GetBareMetal2ProvisionNetworkIpAddressCapacityView, error) {
	var resp view.GetBareMetal2ProvisionNetworkIpAddressCapacityView
	if err := cli.Get("v1/baremetal2/provision-networks/ip-capacity", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

