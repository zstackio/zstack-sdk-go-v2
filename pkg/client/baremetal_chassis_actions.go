// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateBaremetalChassis 更新BaremetalChassis
func (cli *ZSClient) UpdateBaremetalChassis(uuid string, params param.UpdateBaremetalChassisParam) (*view.UpdateBaremetalChassisEventView, error) {
	resp := view.UpdateBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

