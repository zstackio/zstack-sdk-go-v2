// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSlbGroup updates SlbGroup
func (cli *ZSClient) UpdateSlbGroup(uuid string, params param.UpdateSlbGroupParam) (*view.UpdateSlbGroupEventView, error) {
	resp := view.UpdateSlbGroupEventView{}
	if err := cli.Put("v1/load-balancers/slb/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
