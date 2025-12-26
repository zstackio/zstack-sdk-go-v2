// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveLabelFromEventSubscription removes LabelFromEventSubscription
func (cli *ZSClient) RemoveLabelFromEventSubscription(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/events/subscriptions/labels/{uuid}", uuid, string(deleteMode))
}
