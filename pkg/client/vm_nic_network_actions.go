// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeVmNicNetwork 操作VmNicNetwork
func (cli *ZSClient) ChangeVmNicNetwork(params param.ChangeVmNicNetworkParam) (*view.ChangeVmNicNetworkEventView, error) {
	resp := view.ChangeVmNicNetworkEventView{}
	if err := cli.Post("v1/vm-instances/nics/{vmNicUuid}/l3-networks/{destL3NetworkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

