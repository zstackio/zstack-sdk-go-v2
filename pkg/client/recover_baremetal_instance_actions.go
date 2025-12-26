// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RecoverBaremetalInstance operates on BaremetalInstance
func (cli *ZSClient) RecoverBaremetalInstance(uuid string, params param.RecoverBaremetalInstanceParam) (*view.RecoverBaremetalInstanceEventView, error) {
	resp := view.RecoverBaremetalInstanceEventView{}
	if err := cli.Put("v1/baremetal/instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
