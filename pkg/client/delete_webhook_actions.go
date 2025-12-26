// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteWebhook deletes Webhook
func (cli *ZSClient) DeleteWebhook(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/web-hooks/{uuid}", uuid, string(deleteMode))
}
