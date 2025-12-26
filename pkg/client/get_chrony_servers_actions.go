// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetChronyServers gets ChronyServers by uuid
func (cli *ZSClient) GetChronyServers(uuid string) (*view.GetChronyServersView, error) {
	var resp view.GetChronyServersView
	if err := cli.Get("v1/zops/chrony/servers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
