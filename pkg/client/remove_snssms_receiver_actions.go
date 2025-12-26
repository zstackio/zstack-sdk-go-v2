// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveSNSSmsReceiver removes SNSSmsReceiver
func (cli *ZSClient) RemoveSNSSmsReceiver(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/sms-endpoints/{endpointUuid}/receivers/{phoneNumber}", uuid, string(deleteMode))
}
