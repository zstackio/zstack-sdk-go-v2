// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetL3NetworkMtu operates on SetL3NetworkMtu
func (cli *ZSClient) SetL3NetworkMtu(params param.SetL3NetworkMtuParam) (*view.SetL3NetworkMtuEventView, error) {
	resp := view.SetL3NetworkMtuEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/mtu", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
