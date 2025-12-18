// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateBareMetal2ChassisOffering updates BareMetal2ChassisOffering
func (cli *ZSClient) UpdateBareMetal2ChassisOffering(uuid string, params param.UpdateBareMetal2ChassisOfferingParam) (*view.UpdateBareMetal2ChassisOfferingEventView, error) {
	resp := view.UpdateBareMetal2ChassisOfferingEventView{}
	if err := cli.Put("v1/baremetal2/chassis/offerings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
