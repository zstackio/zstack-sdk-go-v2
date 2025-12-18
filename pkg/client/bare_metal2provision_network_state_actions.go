// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeBareMetal2ProvisionNetworkState 操作BareMetal2ProvisionNetworkState
func (cli *ZSClient) ChangeBareMetal2ProvisionNetworkState(uuid string, params param.ChangeBareMetal2ProvisionNetworkStateParam) (*view.ChangeBareMetal2ProvisionNetworkStateEventView, error) {
	resp := view.ChangeBareMetal2ProvisionNetworkStateEventView{}
	if err := cli.Put("v1/baremetal2/provision-networks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

