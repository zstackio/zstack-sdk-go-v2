// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveActionFromEventSubscription 操作RemoveActionFromEventSubscription
func (cli *ZSClient) RemoveActionFromEventSubscription(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/events/subscriptions/{subscriptionUuid}/actions/{actionUuid}", uuid, string(deleteMode))
}

