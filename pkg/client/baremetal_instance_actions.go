// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateBaremetalInstance 创建BaremetalInstance
func (cli *ZSClient) CreateBaremetalInstance(params param.CreateBaremetalInstanceParam) (*view.CreateBaremetalInstanceEventView, error) {
	resp := view.CreateBaremetalInstanceEventView{}
	if err := cli.Post("v1/baremetal/instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

