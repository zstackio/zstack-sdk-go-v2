// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncZBoxCapacity operates on SyncZBoxCapacity
func (cli *ZSClient) SyncZBoxCapacity(uuid string, params param.SyncZBoxCapacityParam) (*view.SyncZBoxCapacityEventView, error) {
	resp := view.SyncZBoxCapacityEventView{}
	if err := cli.Put("v1/zbox/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
