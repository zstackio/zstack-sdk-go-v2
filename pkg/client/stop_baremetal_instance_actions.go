// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// StopBaremetalInstance stops BaremetalInstance
func (cli *ZSClient) StopBaremetalInstance(uuid string, params param.StopBaremetalInstanceParam) (*view.StopBaremetalInstanceEventView, error) {
	resp := view.StopBaremetalInstanceEventView{}
	if err := cli.Put("v1/baremetal/instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
