// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateOvnControllerVm creates OvnControllerVm
func (cli *ZSClient) CreateOvnControllerVm(params param.CreateOvnControllerVmParam) (*view.CreateOvnControllerVmEventView, error) {
	resp := view.CreateOvnControllerVmEventView{}
	if err := cli.Post("v1/ovn/instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
