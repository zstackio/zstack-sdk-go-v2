// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeVmNicNetwork changes VmNicNetwork
func (cli *ZSClient) ChangeVmNicNetwork(uuid string, params param.ChangeVmNicNetworkParam) (*view.ChangeVmNicNetworkEventView, error) {
	resp := view.ChangeVmNicNetworkEventView{}
	if err := cli.Put("v1/vm-instances/nics/{vmNicUuid}/l3-networks/{destL3NetworkUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
