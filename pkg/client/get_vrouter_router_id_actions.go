// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVRouterRouterId gets VRouterRouterId by uuid
func (cli *ZSClient) GetVRouterRouterId(uuid string) (*view.GetVRouterRouterIdView, error) {
	var resp view.GetVRouterRouterIdView
	if err := cli.Get("v1/routerArea/{vRouterUuid}/routerid", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
