// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddBackendServerToServerGroup adds BackendServerToServerGroup
func (cli *ZSClient) AddBackendServerToServerGroup(params param.AddBackendServerToServerGroupParam) (*view.AddBackendServerToServerGroupEventView, error) {
	resp := view.AddBackendServerToServerGroupEventView{}
	if err := cli.Post("v1/load-balancers/servergroups/{serverGroupUuid}/backendservers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
