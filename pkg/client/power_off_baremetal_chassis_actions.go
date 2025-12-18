// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PowerOffBaremetalChassis 操作PowerOffBaremetalChassis
func (cli *ZSClient) PowerOffBaremetalChassis(uuid string, params param.PowerOffBaremetalChassisParam) (*view.PowerOffBaremetalChassisEventView, error) {
	resp := view.PowerOffBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

