// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RebootBaremetalInstance operates on BaremetalInstance
func (cli *ZSClient) RebootBaremetalInstance(uuid string, params param.RebootBaremetalInstanceParam) (*view.RebootBaremetalInstanceEventView, error) {
	resp := view.RebootBaremetalInstanceEventView{}
	if err := cli.Put("v1/baremetal/instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
