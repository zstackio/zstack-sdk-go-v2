// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteCCSCertificate deletes CCSCertificate
func (cli *ZSClient) DeleteCCSCertificate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/crypto/ccs-certificate/delete/{uuid}", uuid, string(deleteMode))
}
