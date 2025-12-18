// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeBaremetalChassisState 操作BaremetalChassisState
func (cli *ZSClient) ChangeBaremetalChassisState(uuid string, params param.ChangeBaremetalChassisStateParam) (*view.ChangeBaremetalChassisStateEventView, error) {
	resp := view.ChangeBaremetalChassisStateEventView{}
	if err := cli.Put("v1/baremetal/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

