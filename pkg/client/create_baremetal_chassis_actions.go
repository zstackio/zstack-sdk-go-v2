// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateBaremetalChassis creates BaremetalChassis
func (cli *ZSClient) CreateBaremetalChassis(params param.CreateBaremetalChassisParam) (*view.CreateBaremetalChassisEventView, error) {
	resp := view.CreateBaremetalChassisEventView{}
	if err := cli.Post("v1/baremetal/chassis", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
