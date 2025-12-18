// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateL2VxlanNetwork creates L2VxlanNetwork
func (cli *ZSClient) CreateL2VxlanNetwork(params param.CreateL2VxlanNetworkParam) (*view.CreateL2VxlanNetworkEventView, error) {
	resp := view.CreateL2VxlanNetworkEventView{}
	if err := cli.Post("v1/l2-networks/vxlan", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
