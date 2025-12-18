// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// InspectBaremetalChassis 操作InspectBaremetalChassis
func (cli *ZSClient) InspectBaremetalChassis(uuid string, params param.InspectBaremetalChassisParam) (*view.InspectBaremetalChassisEventView, error) {
	resp := view.InspectBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

