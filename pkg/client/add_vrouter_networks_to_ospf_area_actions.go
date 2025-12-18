// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddVRouterNetworksToOspfArea adds VRouterNetworksToOspfArea
func (cli *ZSClient) AddVRouterNetworksToOspfArea(params param.AddVRouterNetworksToOspfAreaParam) (*view.AddVRouterNetworksToOspfAreaEventView, error) {
	resp := view.AddVRouterNetworksToOspfAreaEventView{}
	if err := cli.Post("v1/routerArea/{routerAreaUuid}/router/{vRouterUuid}/addnetworks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
