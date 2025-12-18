// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateOvnControllerVm creates OvnControllerVm
func (cli *ZSClient) CreateOvnControllerVm(params param.CreateOvnControllerVmParam) (*view.CreateOvnControllerVmEventView, error) {
	resp := view.CreateOvnControllerVmEventView{}
	if err := cli.Post("v1/ovn/instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
