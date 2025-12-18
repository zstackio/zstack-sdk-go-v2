// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVRouterOspfNeighbor 获取VRouterOspfNeighbor详情
func (cli *ZSClient) GetVRouterOspfNeighbor(uuid string) (*view.GetVRouterOspfNeighborView, error) {
	var resp view.GetVRouterOspfNeighborView
	if err := cli.Get("v1/routerArea/{vRouterUuid}/neighbor", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

