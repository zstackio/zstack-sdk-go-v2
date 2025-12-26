// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncVCenter operates on SyncVCenter
func (cli *ZSClient) SyncVCenter(uuid string, params param.SyncVCenterParam) (*view.SyncVCenterEventView, error) {
	resp := view.SyncVCenterEventView{}
	if err := cli.Put("v1/vcenters/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
