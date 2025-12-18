// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveBackendServerFromServerGroup 操作RemoveBackendServerFromServerGroup
func (cli *ZSClient) RemoveBackendServerFromServerGroup(uuid string, params param.RemoveBackendServerFromServerGroupParam) (*view.RemoveBackendServerFromServerGroupEventView, error) {
	resp := view.RemoveBackendServerFromServerGroupEventView{}
	if err := cli.Put("v1/load-balancers/servergroups/{serverGroupUuid}/backendservers/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

