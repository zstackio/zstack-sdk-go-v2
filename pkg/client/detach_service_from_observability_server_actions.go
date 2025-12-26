// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachServiceFromObservabilityServer operates on ServiceFromObservabilityServer
func (cli *ZSClient) DetachServiceFromObservabilityServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/observability-server/{observabilityServerUuid}/service", uuid, string(deleteMode))
}
