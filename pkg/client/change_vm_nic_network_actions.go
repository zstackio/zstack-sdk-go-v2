// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeVmNicNetwork changes VmNicNetwork
func (cli *ZSClient) ChangeVmNicNetwork(uuid string, params param.ChangeVmNicNetworkParam) (*view.ChangeVmNicNetworkEventView, error) {
	resp := view.ChangeVmNicNetworkEventView{}
	if err := cli.Put("v1/vm-instances/nics/{vmNicUuid}/l3-networks/{destL3NetworkUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
