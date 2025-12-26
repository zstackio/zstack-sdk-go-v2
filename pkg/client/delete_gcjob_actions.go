// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteGCJob deletes GCJob
func (cli *ZSClient) DeleteGCJob(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/gc-jobs/{uuid}", uuid, string(deleteMode))
}
