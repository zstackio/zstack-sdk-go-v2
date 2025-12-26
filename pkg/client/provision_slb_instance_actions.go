// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ProvisionSlbInstance operates on ProvisionSlbInstance
func (cli *ZSClient) ProvisionSlbInstance(uuid string, params param.ProvisionSlbInstanceParam) (*view.ProvisionSlbGroupInstanceEventView, error) {
	resp := view.ProvisionSlbGroupInstanceEventView{}
	if err := cli.Put("v1/load-balancers/slb/instances/{uuid}/provision", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
