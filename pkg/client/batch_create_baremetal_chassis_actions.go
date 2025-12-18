// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// BatchCreateBaremetalChassis 操作BatchCreateBaremetalChassis
func (cli *ZSClient) BatchCreateBaremetalChassis(params param.BatchCreateBaremetalChassisParam) (*view.BatchCreateBaremetalChassisEventView, error) {
	resp := view.BatchCreateBaremetalChassisEventView{}
	if err := cli.Post("v1/baremetal/chassis/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

