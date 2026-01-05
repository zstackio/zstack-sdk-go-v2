// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddL3NetworkToGroup adds L3NetworkToGroup
func (cli *ZSClient) AddL3NetworkToGroup(params param.AddL3NetworkToGroupParam) (*view.AddL3NetworkToGroupEventView, error) {
	resp := view.AddL3NetworkToGroupEventView{}
	if err := cli.Post("v1/nfvinstgroup/group/{nfvInstGroupUuid}/service/{networkServiceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
