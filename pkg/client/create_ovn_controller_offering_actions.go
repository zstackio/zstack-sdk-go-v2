// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateOvnControllerOffering creates OvnControllerOffering
func (cli *ZSClient) CreateOvnControllerOffering(params param.CreateOvnControllerOfferingParam) (*view.CreateInstanceOfferingEventView, error) {
	resp := view.CreateInstanceOfferingEventView{}
	if err := cli.Post("v1/instance-offerings/ovn", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
