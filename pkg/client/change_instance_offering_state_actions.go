// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeInstanceOfferingState changes InstanceOfferingState
func (cli *ZSClient) ChangeInstanceOfferingState(uuid string, params param.ChangeInstanceOfferingStateParam) (*view.ChangeInstanceOfferingStateEventView, error) {
	resp := view.ChangeInstanceOfferingStateEventView{}
	if err := cli.Put("v1/instance-offerings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
