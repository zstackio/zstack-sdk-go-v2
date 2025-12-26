// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveActionFromEventSubscription removes ActionFromEventSubscription
func (cli *ZSClient) RemoveActionFromEventSubscription(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/events/subscriptions/{subscriptionUuid}/actions/{actionUuid}", uuid, string(deleteMode))
}
