// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteCdpPolicy deletes CdpPolicy
func (cli *ZSClient) DeleteCdpPolicy(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cdp-backup-storage/policy/{uuid}", uuid, string(deleteMode))
}
