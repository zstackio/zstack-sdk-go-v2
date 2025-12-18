// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateChronyServers 更新ChronyServers
func (cli *ZSClient) UpdateChronyServers(uuid string, params param.UpdateChronyServersParam) (*view.UpdateChronyServersEventView, error) {
	resp := view.UpdateChronyServersEventView{}
	if err := cli.Put("v1/zops/chrony/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

