// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteNasMountTarget deletes NasMountTarget
func (cli *ZSClient) DeleteNasMountTarget(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/nas/mount/{uuid}", uuid, string(deleteMode))
}
