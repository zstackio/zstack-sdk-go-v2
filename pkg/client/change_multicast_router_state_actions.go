// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeMulticastRouterState changes MulticastRouterState
func (cli *ZSClient) ChangeMulticastRouterState(uuid string, params param.ChangeMulticastRouterStateParam) (*view.ChangeMulticastRouterStateEventView, error) {
	resp := view.ChangeMulticastRouterStateEventView{}
	if err := cli.Put("v1/multicast/virtual-routers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
