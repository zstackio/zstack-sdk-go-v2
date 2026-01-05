// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteNfvInstGroup deletes NfvInstGroup
func (cli *ZSClient) DeleteNfvInstGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/nfvinstgroup/group/{uuid}", uuid, string(deleteMode))
}
