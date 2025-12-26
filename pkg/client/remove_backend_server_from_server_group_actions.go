// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveBackendServerFromServerGroup removes BackendServerFromServerGroup
func (cli *ZSClient) RemoveBackendServerFromServerGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/servergroups/{serverGroupUuid}/backendservers/actions", uuid, string(deleteMode))
}
