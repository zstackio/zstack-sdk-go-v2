// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteSNSTextTemplate deletes SNSTextTemplate
func (cli *ZSClient) DeleteSNSTextTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/alarms/sns/text-templates/{uuid}", uuid, string(deleteMode))
}
