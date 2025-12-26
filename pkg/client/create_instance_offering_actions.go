// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateInstanceOffering creates InstanceOffering
func (cli *ZSClient) CreateInstanceOffering(params param.CreateInstanceOfferingParam) (*view.CreateInstanceOfferingEventView, error) {
	resp := view.CreateInstanceOfferingEventView{}
	if err := cli.Post("v1/instance-offerings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
