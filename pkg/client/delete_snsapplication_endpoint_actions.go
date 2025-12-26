// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteSNSApplicationEndpoint deletes SNSApplicationEndpoint
func (cli *ZSClient) DeleteSNSApplicationEndpoint(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints/{uuid}", uuid, string(deleteMode))
}
