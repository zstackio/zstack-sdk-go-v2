// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateChronyServers updates ChronyServers
func (cli *ZSClient) UpdateChronyServers(uuid string, params param.UpdateChronyServersParam) (*view.UpdateChronyServersEventView, error) {
	resp := view.UpdateChronyServersEventView{}
	if err := cli.Put("v1/zops/chrony/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
