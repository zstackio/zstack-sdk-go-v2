// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVRouterRouterId operates on SetVRouterRouterId
func (cli *ZSClient) SetVRouterRouterId(params param.SetVRouterRouterIdParam) (*view.SetVRouterRouterIdEventView, error) {
	resp := view.SetVRouterRouterIdEventView{}
	if err := cli.Post("v1/routerArea/{vRouterUuid}/routerid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
