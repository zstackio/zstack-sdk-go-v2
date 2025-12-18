// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// InspectBareMetal2Chassis 操作InspectBareMetal2Chassis
func (cli *ZSClient) InspectBareMetal2Chassis(uuid string, params param.InspectBareMetal2ChassisParam) (*view.InspectBareMetal2ChassisEventView, error) {
	resp := view.InspectBareMetal2ChassisEventView{}
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

