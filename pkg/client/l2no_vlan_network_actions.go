// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateL2NoVlanNetwork 创建L2NoVlanNetwork
func (cli *ZSClient) CreateL2NoVlanNetwork(params param.CreateL2NoVlanNetworkParam) (*view.CreateL2NetworkEventView, error) {
	resp := view.CreateL2NetworkEventView{}
	if err := cli.Post("v1/l2-networks/no-vlan", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

