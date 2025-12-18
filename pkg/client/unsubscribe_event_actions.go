// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UnsubscribeEvent operates on UnsubscribeEvent
func (cli *ZSClient) UnsubscribeEvent(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/events/subscriptions/{uuid}", uuid, string(deleteMode))
}
