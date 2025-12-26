// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncChronyServers operates on SyncChronyServers
func (cli *ZSClient) SyncChronyServers(uuid string, params param.SyncChronyServersParam) (*view.SyncChronyServersEventView, error) {
	resp := view.SyncChronyServersEventView{}
	if err := cli.Put("v1/zops/chrony/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
