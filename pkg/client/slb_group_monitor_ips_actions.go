// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeSlbGroupMonitorIps 操作SlbGroupMonitorIps
func (cli *ZSClient) ChangeSlbGroupMonitorIps(uuid string, params param.ChangeSlbGroupMonitorIpsParam) (*view.ChangeSlbGroupMonitorIpsEventView, error) {
	resp := view.ChangeSlbGroupMonitorIpsEventView{}
	if err := cli.Put("v1/load-balancers/slb/groups/{slbGroupUuid}/monitorIps", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

