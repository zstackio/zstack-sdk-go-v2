// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateBareMetal2ProvisionNetwork creates BareMetal2ProvisionNetwork
func (cli *ZSClient) CreateBareMetal2ProvisionNetwork(params param.CreateBareMetal2ProvisionNetworkParam) (*view.CreateBareMetal2ProvisionNetworkEventView, error) {
	resp := view.CreateBareMetal2ProvisionNetworkEventView{}
	if err := cli.Post("v1/baremetal2/provision-networks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
