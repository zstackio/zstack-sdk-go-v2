// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeVpcHaGroupMonitorIps changes VpcHaGroupMonitorIps
func (cli *ZSClient) ChangeVpcHaGroupMonitorIps(uuid string, params param.ChangeVpcHaGroupMonitorIpsParam) (*view.ChangeVpcHaGroupMonitorIpsEventView, error) {
	resp := view.ChangeVpcHaGroupMonitorIpsEventView{}
	if err := cli.Put("v1/vpc/hagroups/{uuid}/monitorIps", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
