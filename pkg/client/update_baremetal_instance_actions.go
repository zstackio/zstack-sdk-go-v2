// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateBaremetalInstance updates BaremetalInstance
func (cli *ZSClient) UpdateBaremetalInstance(uuid string, params param.UpdateBaremetalInstanceParam) (*view.UpdateBaremetalInstanceEventView, error) {
	resp := view.UpdateBaremetalInstanceEventView{}
	if err := cli.Put("v1/baremetal/instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
