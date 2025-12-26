// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAliyunNasAccessGroup deletes AliyunNasAccessGroup
func (cli *ZSClient) DeleteAliyunNasAccessGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/nas/access/{uuid}", uuid, string(deleteMode))
}
