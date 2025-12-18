// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateOvnControllerOffering 创建OvnControllerOffering
func (cli *ZSClient) CreateOvnControllerOffering(params param.CreateOvnControllerOfferingParam) (*view.CreateInstanceOfferingEventView, error) {
	resp := view.CreateInstanceOfferingEventView{}
	if err := cli.Post("v1/instance-offerings/ovn", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

