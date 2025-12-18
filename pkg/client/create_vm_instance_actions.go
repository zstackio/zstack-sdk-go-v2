// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVmInstance creates VmInstance
func (cli *ZSClient) CreateVmInstance(params param.CreateVmInstanceParam) (*view.CreateVmInstanceEventView, error) {
	resp := view.CreateVmInstanceEventView{}
	if err := cli.Post("v1/vm-instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
