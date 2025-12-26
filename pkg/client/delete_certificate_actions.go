// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteCertificate deletes Certificate
func (cli *ZSClient) DeleteCertificate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/certificates/{uuid}", uuid, string(deleteMode))
}
