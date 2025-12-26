// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddHostRouteToL3Network adds HostRouteToL3Network
func (cli *ZSClient) AddHostRouteToL3Network(params param.AddHostRouteToL3NetworkParam) (*view.AddHostRouteToL3NetworkEventView, error) {
	resp := view.AddHostRouteToL3NetworkEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/hostroute", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
