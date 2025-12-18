// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryBareMetal2ChassisOffering 查询BareMetal2ChassisOffering列表
func (cli *ZSClient) QueryBareMetal2ChassisOffering(params param.QueryParam) ([]view.QueryBareMetal2ChassisOfferingView, error) {
	var resp []view.QueryBareMetal2ChassisOfferingView
	return resp, cli.List("v1/baremetal2/chassis/offerings", &params, &resp)
}

