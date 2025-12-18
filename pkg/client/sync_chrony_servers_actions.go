// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncChronyServers 操作SyncChronyServers
func (cli *ZSClient) SyncChronyServers(uuid string, params param.SyncChronyServersParam) (*view.SyncChronyServersEventView, error) {
	resp := view.SyncChronyServersEventView{}
	if err := cli.Put("v1/zops/chrony/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

