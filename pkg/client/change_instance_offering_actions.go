// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeInstanceOffering changes InstanceOffering
func (cli *ZSClient) ChangeInstanceOffering(uuid string, params param.ChangeInstanceOfferingParam) (*view.ChangeInstanceOfferingEventView, error) {
	resp := view.ChangeInstanceOfferingEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
