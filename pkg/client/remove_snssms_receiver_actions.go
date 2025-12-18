// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveSNSSmsReceiver 操作RemoveSNSSmsReceiver
func (cli *ZSClient) RemoveSNSSmsReceiver(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/sms-endpoints/{endpointUuid}/receivers/{phoneNumber}", uuid, string(deleteMode))
}

