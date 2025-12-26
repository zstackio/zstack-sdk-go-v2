// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateInstanceOffering updates InstanceOffering
func (cli *ZSClient) UpdateInstanceOffering(uuid string, params param.UpdateInstanceOfferingParam) (*view.UpdateInstanceOfferingEventView, error) {
	resp := view.UpdateInstanceOfferingEventView{}
	if err := cli.Put("v1/instance-offerings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
