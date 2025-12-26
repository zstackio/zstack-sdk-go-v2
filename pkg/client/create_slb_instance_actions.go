// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSlbInstance creates SlbInstance
func (cli *ZSClient) CreateSlbInstance(params param.CreateSlbInstanceParam) (*view.CreateSlbInstanceEventView, error) {
	resp := view.CreateSlbInstanceEventView{}
	if err := cli.Post("v1/load-balancers/slb/instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
