// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PowerOnBaremetalChassis 操作PowerOnBaremetalChassis
func (cli *ZSClient) PowerOnBaremetalChassis(uuid string, params param.PowerOnBaremetalChassisParam) (*view.PowerOnBaremetalChassisEventView, error) {
	resp := view.PowerOnBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

