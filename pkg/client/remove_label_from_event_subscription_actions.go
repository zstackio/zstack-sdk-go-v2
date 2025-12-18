// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveLabelFromEventSubscription 操作RemoveLabelFromEventSubscription
func (cli *ZSClient) RemoveLabelFromEventSubscription(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/events/subscriptions/labels/{uuid}", uuid, string(deleteMode))
}

