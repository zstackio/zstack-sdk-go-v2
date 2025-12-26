// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVmInstanceFromOvf creates VmInstanceFromOvf
func (cli *ZSClient) CreateVmInstanceFromOvf(params param.CreateVmInstanceFromOvfParam) (*view.CreateVmInstanceFromOvfEventView, error) {
	resp := view.CreateVmInstanceFromOvfEventView{}
	if err := cli.Post("v1/ovf/create-vm-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
