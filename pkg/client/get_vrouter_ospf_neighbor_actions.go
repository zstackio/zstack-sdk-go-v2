// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVRouterOspfNeighbor gets VRouterOspfNeighbor by uuid
func (cli *ZSClient) GetVRouterOspfNeighbor(uuid string) (*view.GetVRouterOspfNeighborView, error) {
	var resp view.GetVRouterOspfNeighborView
	if err := cli.Get("v1/routerArea/{vRouterUuid}/neighbor", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
