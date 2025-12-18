// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeBareMetal2ChassisOfferingState 操作BareMetal2ChassisOfferingState
func (cli *ZSClient) ChangeBareMetal2ChassisOfferingState(uuid string, params param.ChangeBareMetal2ChassisOfferingStateParam) (*view.ChangeBareMetal2ChassisOfferingStateEventView, error) {
	resp := view.ChangeBareMetal2ChassisOfferingStateEventView{}
	if err := cli.Put("v1/baremetal2/chassis/offerings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

