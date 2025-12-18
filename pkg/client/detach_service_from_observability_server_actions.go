// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachServiceFromObservabilityServer operates on ServiceFromObservabilityServer
func (cli *ZSClient) DetachServiceFromObservabilityServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/observability-server/{observabilityServerUuid}/service", uuid, string(deleteMode))
}
