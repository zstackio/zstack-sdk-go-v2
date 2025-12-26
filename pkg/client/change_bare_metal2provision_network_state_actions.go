// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeBareMetal2ProvisionNetworkState changes BareMetal2ProvisionNetworkState
func (cli *ZSClient) ChangeBareMetal2ProvisionNetworkState(uuid string, params param.ChangeBareMetal2ProvisionNetworkStateParam) (*view.ChangeBareMetal2ProvisionNetworkStateEventView, error) {
	resp := view.ChangeBareMetal2ProvisionNetworkStateEventView{}
	if err := cli.Put("v1/baremetal2/provision-networks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
