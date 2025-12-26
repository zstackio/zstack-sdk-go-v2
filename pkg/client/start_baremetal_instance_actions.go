// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// StartBaremetalInstance starts BaremetalInstance
func (cli *ZSClient) StartBaremetalInstance(uuid string, params param.StartBaremetalInstanceParam) (*view.StartBaremetalInstanceEventView, error) {
	resp := view.StartBaremetalInstanceEventView{}
	if err := cli.Put("v1/baremetal/instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
