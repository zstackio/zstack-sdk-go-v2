// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateL2VxlanNetworkPool creates L2VxlanNetworkPool
func (cli *ZSClient) CreateL2VxlanNetworkPool(params param.CreateL2VxlanNetworkPoolParam) (*view.CreateL2VxlanNetworkPoolEventView, error) {
	resp := view.CreateL2VxlanNetworkPoolEventView{}
	if err := cli.Post("v1/l2-networks/vxlan-pool", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
