// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateBaremetalInstance creates BaremetalInstance
func (cli *ZSClient) CreateBaremetalInstance(params param.CreateBaremetalInstanceParam) (*view.CreateBaremetalInstanceEventView, error) {
	resp := view.CreateBaremetalInstanceEventView{}
	if err := cli.Post("v1/baremetal/instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
