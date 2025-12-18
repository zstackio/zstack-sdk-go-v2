// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteWebhook deletes Webhook
func (cli *ZSClient) DeleteWebhook(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/web-hooks/{uuid}", uuid, string(deleteMode))
}
