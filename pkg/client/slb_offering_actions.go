// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSlbOffering 创建SlbOffering
func (cli *ZSClient) CreateSlbOffering(params param.CreateSlbOfferingParam) (*view.CreateInstanceOfferingEventView, error) {
	resp := view.CreateInstanceOfferingEventView{}
	if err := cli.Post("v1/instance-offerings/slb", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

