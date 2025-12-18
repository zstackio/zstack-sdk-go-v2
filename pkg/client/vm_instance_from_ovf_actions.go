// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVmInstanceFromOvf 创建VmInstanceFromOvf
func (cli *ZSClient) CreateVmInstanceFromOvf(params param.CreateVmInstanceFromOvfParam) (*view.CreateVmInstanceFromOvfEventView, error) {
	resp := view.CreateVmInstanceFromOvfEventView{}
	if err := cli.Post("v1/ovf/create-vm-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

