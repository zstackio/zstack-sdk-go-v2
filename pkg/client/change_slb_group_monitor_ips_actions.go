// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeSlbGroupMonitorIps changes SlbGroupMonitorIps
func (cli *ZSClient) ChangeSlbGroupMonitorIps(uuid string, params param.ChangeSlbGroupMonitorIpsParam) (*view.ChangeSlbGroupMonitorIpsEventView, error) {
	resp := view.ChangeSlbGroupMonitorIpsEventView{}
	if err := cli.Put("v1/load-balancers/slb/groups/{slbGroupUuid}/monitorIps", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
