// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAliyunNasAccessGroupRule deletes AliyunNasAccessGroupRule
func (cli *ZSClient) DeleteAliyunNasAccessGroupRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/nas/access/rule/{uuid}", uuid, string(deleteMode))
}
