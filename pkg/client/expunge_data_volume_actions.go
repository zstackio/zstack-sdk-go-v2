// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// ExpungeDataVolume operates on DataVolume
func (cli *ZSClient) ExpungeDataVolume(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volumes/{uuid}/actions", uuid, string(deleteMode))
}
